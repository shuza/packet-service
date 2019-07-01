package main

import (
	"context"
	"errors"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	boxPb "github.com/shuza/box-service/proto"
	"github.com/shuza/packet-service/db"
	pb "github.com/shuza/packet-service/proto"
	"github.com/shuza/packet-service/service"
	userPb "github.com/shuza/porter/user-service/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"os"
)

var (
	srv micro.Service
)

func main() {
	repo := &db.MongoRepository{}
	mongoUri := os.Getenv("MONGO_HOST")
	if err := repo.Init(mongoUri); err != nil {
		panic(err)
	}
	defer repo.Close()

	srv := micro.NewService(
		micro.Name("porter.packet"),
		micro.Version("latest"),
		micro.WrapHandler(AuthWrapper),
	)
	boxClient := boxPb.NewBoxServiceClient(os.Getenv("BOX_SERVICE_ADDRESS"), srv.Client())

	//	It will parse the command line flags
	srv.Init()

	packetService := service.NewPacketService(repo, boxClient)
	pb.RegisterPacketServiceHandler(srv.Server(), &packetService)

	log.Println("user service is running...")
	if err := srv.Run(); err != nil {
		panic(err)
	}

}

func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		meta, ok := metadata.FromIncomingContext(ctx)
		if !ok || len(meta["token"]) < 1 {
			return errors.New("no auth meta-data found in request")
		}

		token := meta["token"][0]
		log.Println("token :  %v\n", token)

		//	Authenticate request here
		authClient := userPb.NewUserServiceClient(os.Getenv("USER_SERVICE_ADDRESS"), srv.Client())
		if _, err := authClient.ValidateToken(ctx, &userPb.Token{Token: token}); err != nil {
			return err
		}

		err := fn(ctx, req, resp)
		return err

	}
}

func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Warnf("Missing context metadata")
		return nil, errors.New("missing context metadata")
	}
	if len(meta["token"]) != 1 {
		log.Warnf("invalid token len != 1")
		return nil, errors.New("invalid token")
	}

	log.Println("Token  ==   ", meta["token"][0])
	return handler(ctx, req)
}

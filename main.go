package main

import (
	boxPb "github.com/shuza/box-service/proto"
	"github.com/shuza/packet-service/db"
	pb "github.com/shuza/packet-service/proto"
	"github.com/shuza/packet-service/service"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
)

func main() {
	repo := &db.Repository{}
	port := os.Getenv("PORT")

	//	setup gRPC server
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln("failed to listen  :  ", err)
	}

	//	box service connection
	conn, err := grpc.Dial(os.Getenv("BOX_SERVICE_ADDRESS"), grpc.WithInsecure())
	if err != nil {
		log.Fatalln("failed to connect Box Service :  ", err)
	}
	boxClient := boxPb.NewBoxServiceClient(conn)

	s := grpc.NewServer()

	//	Register our service with gRPC server
	//	this will tie our implementation into the auto-generated interface code
	//	for our protobuf edition
	packetService := service.NewPacketService(repo, boxClient)
	pb.RegisterPacketServiceServer(s, &packetService)

	// Register reflection service on gRPC server.
	reflection.Register(s)

	log.Println("Running on port :  ", port)
	if err := s.Serve(listen); err != nil {
		log.Fatalln("failed to server  :  ", err)
	}
}

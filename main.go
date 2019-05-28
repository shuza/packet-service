package main

import (
	"context"
	"github.com/joho/godotenv"
	pb "github.com/shuza/packet-service/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"sync"
)

type repository interface {
	Create(*pb.Packet) (*pb.Packet, error)
	GetAll() []*pb.Packet
}

type Repository struct {
	mu      sync.RWMutex
	packets []*pb.Packet
}

//	Create new packet
func (repo *Repository) Create(packet *pb.Packet) (*pb.Packet, error) {
	repo.mu.Lock()
	updated := append(repo.packets, packet)
	repo.packets = updated
	repo.mu.Unlock()
	return packet, nil
}

//	Get all packets
func (repo *Repository) GetAll() []*pb.Packet {
	return repo.packets
}

type service struct {
	repo repository
}

func (s *service) CreatePacket(ctx context.Context, req *pb.Packet) (*pb.Response, error) {
	packet, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}

	return &pb.Response{Created: true, Packet: packet}, nil
}

func (s *service) GetPackets(ctx context.Context, req *pb.Empty) (*pb.Response, error) {
	packets := s.repo.GetAll()
	return &pb.Response{Packets: packets}, nil
}

func main() {
	godotenv.Load()
	repo := &Repository{}
	port := os.Getenv("PORT")

	//	setup gRPC server
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln("failed to listen  :  ", err)
	}

	s := grpc.NewServer()

	//	Register our service with gRPC server
	//	this will tie our implementation into the auto-generated interface code
	//	for our protobuf edition
	pb.RegisterPacketServiceServer(s, &service{repo})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	log.Println("Running on port :  ", port)
	if err := s.Serve(listen); err != nil {
		log.Fatalln("failed to server  :  ", err)
	}
}

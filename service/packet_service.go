package service

import (
	"context"
	boxPb "github.com/shuza/box-service/proto"
	"github.com/shuza/packet-service/db"
	pb "github.com/shuza/packet-service/proto"
	"log"
)

/**
 *  := 	create date: 31-May-2019
 *  := 	(C) CopyRight Shuza
 *  := 	shuza.ninja
 *  := 	shuza.sa@gmail.com
 *  := 	Code	:	Coffee	: Fun
 **/

type service struct {
	repo             db.IRepository
	boxServiceClient boxPb.BoxServiceClient
}

func NewPacketService(repo db.IRepository, boxClient boxPb.BoxServiceClient) service {
	return service{repo: repo, boxServiceClient: boxClient}
}

func (s *service) CreatePacket(ctx context.Context, req *pb.Packet) (*pb.Response, error) {
	//	Call box service to find available box with packet weight
	boxResponse, err := s.boxServiceClient.FindAvailableBox(context.Background(), &boxPb.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Items)),
	})

	if err != nil {
		return nil, err
	}

	log.Println("Found box :  ", boxResponse.Box.Id)
	req.BoxId = boxResponse.Box.Id

	err = s.repo.Create(req)
	if err != nil {
		return nil, err
	}

	return &pb.Response{Created: true}, nil
}

func (s *service) GetPackets(ctx context.Context, req *pb.Empty) (*pb.Response, error) {
	packets, err := s.repo.GetAll()
	return &pb.Response{Packets: packets}, err
}

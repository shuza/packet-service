package service

import (
	"context"
	"github.com/shuza/packet-service/db"
	pb "github.com/shuza/packet-service/proto"
)

/**
 *  := 	create date: 31-May-2019
 *  := 	(C) CopyRight Shuza
 *  := 	shuza.ninja
 *  := 	shuza.sa@gmail.com
 *  := 	Code	:	Coffee	: Fun
 **/

type service struct {
	repo db.IRepository
}

func NewPacketService(repo db.IRepository) service {
	return service{repo: repo}
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

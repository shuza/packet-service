package db

import (
	pb "github.com/shuza/packet-service/proto"
	"github.com/stretchr/testify/mock"
	"sync"
)

/**
 *  := 	create date: 31-May-2019
 *  := 	(C) CopyRight Shuza
 *  := 	shuza.ninja
 *  := 	shuza.sa@gmail.com
 *  := 	Code	:	Coffee	: Fun
 **/

type MockDb struct {
	mu      sync.RWMutex
	packets []*pb.Packet
	mock.Mock
}

//	Create new packet
func (repo *MockDb) Create(packet *pb.Packet) (*pb.Packet, error) {
	args := repo.Mock.Called(packet)
	if args.Get(1) != nil {
		return args.Get(0).(*pb.Packet), args.Get(1).(error)
	}

	repo.mu.Lock()
	updated := append(repo.packets, packet)
	repo.packets = updated
	repo.mu.Unlock()
	return args.Get(0).(*pb.Packet), nil
}

//	Get all packets
func (repo *MockDb) GetAll() []*pb.Packet {
	return repo.packets
}

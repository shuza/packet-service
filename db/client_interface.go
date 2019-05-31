package db

import (
	pb "github.com/shuza/packet-service/proto"
	"sync"
)

/**
 *  := 	create date: 31-May-2019
 *  := 	(C) CopyRight Shuza
 *  := 	shuza.ninja
 *  := 	shuza.sa@gmail.com
 *  := 	Code	:	Coffee	: Fun
 **/

type IRepository interface {
	Create(*pb.Packet) (*pb.Packet, error)
	GetAll() []*pb.Packet
}

var client IRepository

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

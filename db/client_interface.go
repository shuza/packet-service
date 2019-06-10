package db

import (
	pb "github.com/shuza/packet-service/proto"
)

/**
 *  := 	create date: 31-May-2019
 *  := 	(C) CopyRight Shuza
 *  := 	shuza.ninja
 *  := 	shuza.sa@gmail.com
 *  := 	Code	:	Coffee	: Fun
 **/

type IRepository interface {
	Init(host string) error
	Create(*pb.Packet) error
	GetAll() ([]*pb.Packet, error)
	Close()
}

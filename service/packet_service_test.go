package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/shuza/packet-service/db"
	pb "github.com/shuza/packet-service/proto"
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"log"
	"testing"
)

/**
 *  := 	create date: 31-May-2019
 *  := 	(C) CopyRight Shuza
 *  := 	shuza.ninja
 *  := 	shuza.sa@gmail.com
 *  := 	Code	:	Coffee	: Fun
 **/

func TestService_CreatePacket(t *testing.T) {
	packet, err := parseFile("./../packet-cli/packet.json")
	if err != nil {
		log.Panic(err)
	}

	repo := &db.MockDb{}
	repo.Mock.On("Create", packet).Return(packet, nil)
	packetService := NewPacketService(repo)
	Convey("Requesting to create a new valid packet", t, func() {
		resp, err := packetService.CreatePacket(context.Background(), packet)
		Convey("Check response and error", func() {
			So(err, ShouldBeNil)
			So(resp.Created, ShouldBeTrue)
		})
	})

	repo = &db.MockDb{}
	repo.Mock.On("Create", packet).Return(packet, errors.New("Can't create packet"))
	packetService = NewPacketService(repo)
	Convey("Requesting to create an invalid packet", t, func() {
		resp, err := packetService.CreatePacket(context.Background(), packet)
		Convey("Should fail the operation", func() {
			So(err, ShouldNotBeNil)
			So(resp, ShouldBeNil)
		})
	})
}

func TestService_GetPackets(t *testing.T) {
	packet, err := parseFile("./../packet-cli/packet.json")
	if err != nil {
		log.Panic(err)
	}

	repo := &db.MockDb{}
	repo.Mock.On("Create", packet).Return(packet, nil)
	packetService := NewPacketService(repo)
	Convey("Requesting for empty packet list", t, func() {
		resp, err := packetService.GetPackets(context.Background(), &pb.Empty{})
		Convey("It should return empty packet list", func() {
			So(err, ShouldBeNil)
			So(len(resp.Packets), ShouldBeZeroValue)
		})
	})

	_, _ = packetService.CreatePacket(context.Background(), packet)
	Convey("Requesting for packet list", t, func() {
		resp, err := packetService.GetPackets(context.Background(), &pb.Empty{})
		Convey("It should return one packet in the list", func() {
			So(err, ShouldBeNil)
			So(len(resp.Packets), ShouldEqual, 1)
		})
	})
}

func parseFile(file string) (*pb.Packet, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var packet *pb.Packet
	err = json.Unmarshal(data, &packet)

	return packet, err
}

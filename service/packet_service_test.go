package service

import (
	"context"
	"encoding/json"
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
	Convey("Requesting to create a new packet", t, func() {
		resp, err := packetService.CreatePacket(context.Background(), packet)
		Convey("Check response and error", func() {
			So(err, ShouldBeNil)
			So(resp.Created, ShouldBeTrue)
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

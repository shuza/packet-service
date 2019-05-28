package main

import (
	"context"
	"encoding/json"
	pb "github.com/shuza/packet-service/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"io/ioutil"
	"os"
)

const (
	address         = "localhost:8000"
	defaultFilename = "packet.json"
)

func parseFile(file string) (*pb.Packet, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var packet *pb.Packet
	err = json.Unmarshal(data, &packet)

	return packet, err
}

func main() {
	//	setup connection to the gRPC server
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Can't connect to server  :  %v", err)
	}
	defer conn.Close()

	client := pb.NewPacketServiceClient(conn)
	filename := defaultFilename
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	packet, err := parseFile(filename)
	if err != nil {
		log.Fatalln("Can't parse file  :  ", err)
	}

	response, err := client.CreatePacket(context.Background(), packet)
	if err != nil {
		log.Fatalln("Can't create packet  :  ", err)
	}
	log.Println("Created  :  ", response.Created)

	packets, err := client.GetPackets(context.Background(), &pb.Empty{})
	for _, v := range packets.GetPackets() {
		log.Println(v)
	}
}

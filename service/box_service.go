package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	boxPb "github.com/shuza/box-service/proto"
	pb "github.com/shuza/packet-service/proto"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"packet-service/model"
)

/**
 *  := 	create date: 31-May-2019
 *  := 	(C) CopyRight Shuza
 *  := 	shuza.ninja
 *  := 	shuza.sa@gmail.com
 *  := 	Code	:	Coffee	: Fun
 **/

type service struct {
	client        http.Client
	boxServiceUrl string
}

func NewPacketService(client http.Client, serviceUrl string) service {
	return service{client: client, boxServiceUrl: serviceUrl}
}

func (s service) FindSuitableBox(capacity int32, weight int32) (model.Box, error) {
	url := fmt.Sprintf("%s/box", s.boxServiceUrl)
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.URL.Query().Add("capacity", string(capacity))
	req.URL.Query().Add("capacity", string(weight))
	req.URL.RawQuery = req.URL.Query().Encode()

	res, err := s.client.Do(req)
	if err != nil {
		log.Warnf("Can't connect to find suitable box service Error :  %v\n", err)
		return model.Box{}, err
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	response := make(map[string]interface{})
	json.Unmarshal(body, &response)
	if response["status"] != http.StatusOK {
		log.Warnf("find box service failed. Response :  %v\n", response)
		return model.Box{}, errors.New(fmt.Sprintf("Can't find suitable box"))
	}

	var box model.Box
	if err:=json.Unmarshal(response["data"], &box); err!=nil{

	}
}

func (s *service) CreatePacket(ctx context.Context, req *pb.Packet, resp *pb.Response) error {
	//	Call box service to find available box with packet weight
	boxResponse, err := s.boxServiceClient.FindAvailableBox(context.Background(), &boxPb.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Items)),
	})

	if err != nil {
		return err
	}

	log.Println("Found box :  ", boxResponse.Box.Id)
	req.BoxId = boxResponse.Box.Id

	err = s.repo.Create(req)
	if err != nil {
		return err
	}
	resp.Created = true
	resp.Packet = req
	return nil
}

func (s *service) GetPackets(ctx context.Context, req *pb.Empty, resp *pb.Response) error {
	packets, err := s.repo.GetAll()
	if err != nil {
		return err
	}
	resp.Packets = packets
	return nil
}

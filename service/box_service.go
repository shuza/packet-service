package service

import (
	"encoding/json"
	"errors"
	"fmt"
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

type boxService struct {
	client        http.Client
	boxServiceUrl string
}

func NewBoxService(client http.Client, serviceUrl string) boxService {
	return boxService{client: client, boxServiceUrl: serviceUrl}
}

func (s boxService) FindSuitableBox(capacity int32, weight int32) (model.Box, error) {
	url := fmt.Sprintf("%s/box", s.boxServiceUrl)
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	q := req.URL.Query()

	fmt.Printf("capacity : %d    weight : %d\n", capacity, weight)
	q.Add("capacity", fmt.Sprintf("%d", capacity))
	q.Add("max_weight", fmt.Sprintf("%d", weight))
	req.URL.RawQuery = q.Encode()

	res, err := s.client.Do(req)
	if err != nil || res.StatusCode != http.StatusOK {
		log.Warnf("Can't connect to find suitable box boxService Error :  %v\n", err)
		return model.Box{}, errors.New(fmt.Sprintf("box service status %v Error : %v", res.Status, err))
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	var box model.Box
	if err := json.Unmarshal(body, &box); err != nil {
		log.Warnf("Can't parse find box boxService response Error :  %v\n", err)
		return model.Box{}, err
	}

	return box, nil
}

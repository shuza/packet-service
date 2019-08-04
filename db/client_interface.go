package db

import (
	"packet-service/model"
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
	Create(packet model.Packet) error
	GetAll() ([]model.Packet, error)
	Close()
}

var Client IRepository

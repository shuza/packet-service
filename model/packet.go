package model

import "gopkg.in/mgo.v2/bson"

type Packet struct {
	Id          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Description string        `json:"description"`
	Weight      int32         `json:"weight"`
	BoxId       bson.ObjectId `json:"box_id"`
	Items       []Item        `json:"items"`
}

type Item struct {
	Id          string `json:"id"`
	CustomerId  string `json:"customer_id"`
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
}

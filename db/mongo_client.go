package db

import (
	"gopkg.in/mgo.v2"
	"packet-service/model"
)

type MongoRepository struct {
	session *mgo.Session
}

func (repo *MongoRepository) Init(host string) error {
	session, err := mgo.Dial(host)
	if err != nil {
		return err
	}

	session.SetMode(mgo.Monotonic, true)
	repo.session = session

	return nil
}

//	Create new packet
func (repo *MongoRepository) Create(packet model.Packet) error {
	return repo.collection().Insert(packet)
}

//	Get all packets
func (repo *MongoRepository) GetAll() ([]model.Packet, error) {
	packets := make([]model.Packet, 0)
	err := repo.collection().Find(nil).All(&packets)

	return packets, err
}

func (repo *MongoRepository) collection() *mgo.Collection {
	return repo.session.DB("porter").C("packets")
}

func (repo *MongoRepository) Close() {
	repo.session.Close()
}

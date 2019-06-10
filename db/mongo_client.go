package db

import (
	pb "github.com/shuza/packet-service/proto"
	"gopkg.in/mgo.v2"
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
func (repo *MongoRepository) Create(packet *pb.Packet) error {
	return repo.collection().Insert(packet)
}

//	Get all packets
func (repo *MongoRepository) GetAll() ([]*pb.Packet, error) {
	packets := make([]*pb.Packet, 0)
	err := repo.collection().Find(nil).All(&packets)

	return packets, err
}

func (repo *MongoRepository) collection() *mgo.Collection {
	return repo.session.DB("porter").C("packets")
}

func (repo *MongoRepository) Close() {
	repo.session.Close()
}

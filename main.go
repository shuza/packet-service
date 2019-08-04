package main

import (
	"fmt"
	"os"
	"packet-service/api"
	"packet-service/db"
)

func main() {
	initDB()

	r := api.NewGinEngine()

	fmt.Println("Packet service is running on port 8082")
	r.Run(":8082")
}

func initDB() {
	db.Client = &db.MongoRepository{}
	if err := db.Client.Init(os.Getenv("MONGO_HOST")); err != nil {
		panic(err)
	}

	r := api.NewGinEngine()
	fmt.Println("Packet service is running on port :8082")
	if err := r.Run(":8082"); err != nil {
		panic(err)
	}
}

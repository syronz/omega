package main

import (
	"rest-gin-gorm/initiate"
	"rest-gin-gorm/server"
)

func main() {

	cfg := initiate.Setup()
	defer cfg.DB.Close()

	s := server.Setup(cfg)

	err := s.Run()
	if err != nil {
		cfg.Log.Fatal(err)
	}
}

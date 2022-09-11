package app

import (
	"log"
	"taskmanager/internal/db"
	"taskmanager/internal/models"
)

func Start() {
	// config init
	var cfg models.Config
	err := cfg.Init("./config.yml")
	if err != nil {
		log.Fatal(err)
	}

	// storage init
	var str db.DB
	err = str.Init(cfg.DBPath)
	if err != nil {
		log.Fatal(err)
	}

	s := models.NewServer(
		models.WithConfig(cfg),
		models.WithStorage(&str))

	s.SetupApp()

	log.Println(s.Listen())
}

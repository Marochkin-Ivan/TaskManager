package models

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

func (s *Server) SetupApp() {
	a := fiber.New()

	a.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))
	a.Use(recover.New(recover.Config{EnableStackTrace: true}))
	a.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path} ${resBody}\n",
	}))

	// groups
	api := s.App.Group("/api")
	user := api.Group("/user")
	board := api.Group("/board")
	task := board.Group("/task")

	// ____start HANDLERS____
	user.Post("registration", s.Registration)
	user.Post("authorization", s.Authorization)

	// ____end HANDLERS____

	s.App = a
}

func (s Server) Listen() error {
	return s.App.Listen(s.Cfg.Host)
}

//func (c *Config) Init() {
//	c.Host = "localhost:8080"
//}

func (c *Config) Init(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	d := yaml.NewDecoder(file)

	if err := d.Decode(&c); err != nil {
		return err
	}

	return nil
}

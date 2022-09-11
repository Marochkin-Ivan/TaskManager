package models

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type Server struct {
	App     *fiber.App
	Cfg     Config
	Storage Storable4Server
}

type Config struct {
	Host   string `yaml:"host"`
	DBPath string `yaml:"db_path"`
}

type User struct {
	Login    string `json:"login" form:"login"`
	Password string `json:"password" form:"password"`
}

type UserResponse struct {
	Token string `json:"token" form:"token"`
}

type Board struct {
	Name    string
	Owner   string
	Members []User
	Tasks   []Task
}

type Task struct {
	Name        string
	Description string
	Deadline    string
	Status      string
	Importance  string
}

type Storable4User interface {
	RegistrationUser(u User) (string, bool, error)
	AuthorizationUser(u User) (string, bool, error)
}

type Storable4Server interface {
	Storable4User
}

type Claims struct {
	jwt.StandardClaims
	ID string `json:"id"`
}

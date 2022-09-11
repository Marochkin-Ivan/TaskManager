package models

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (s *Server) Registration(c *fiber.Ctx) error {
	var u User
	err := c.BodyParser(&u)
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	token, exist, err := s.Storage.RegistrationUser(u)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	if exist {
		return c.SendStatus(http.StatusConflict)
	}

	jwt := GenerateToken(token)
	return c.Status(http.StatusOK).JSON(UserResponse{Token: jwt})
}

func (s *Server) Authorization(c *fiber.Ctx) error {
	var u User
	err := c.BodyParser(&u)
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	uid, ok, err := s.Storage.AuthorizationUser(u)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}
	if !ok {
		return c.SendStatus(http.StatusUnauthorized)
	}

	jwt := GenerateToken(uid)
	return c.Status(http.StatusOK).JSON(UserResponse{Token: jwt})
}

package http

import (
	"context"
	"net/http"
	"simple-redis-go/internal/application"
	"simple-redis-go/internal/domain"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Biz interface {
	CreateUser(string, string, string, string) (*domain.User, error)
	Login(string, string) (string, error)
	GetCountUser() (int64, error)
	GetPingCount(context.Context, int64) (int, error)
	GetTopUser(int) ([]domain.UserItemTopUser, error)
	Ping(string) (int, error)
	GetUserFromSession(string) (*domain.User, error)
}

type Api struct {
	biz Biz
}

func NewApi(db *gorm.DB, rdb *redis.Client) Api {
	biz := application.NewAuthServer(db, rdb)
	return Api{
		biz: biz,
	}
}

type CreateUserBody struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (api *Api) Register(c *fiber.Ctx) error {
	var body CreateUserBody
	if err := c.BodyParser(&body); err != nil {
		c.Status(http.StatusBadRequest)
		return err
	}
	_, err := api.biz.CreateUser(body.Email, body.Password, body.FirstName, body.LastName)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return err
	}
	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"data": "OK",
	})
}

type LoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (api *Api) Login(c *fiber.Ctx) error {
	body := new(LoginBody)
	if err := c.BodyParser(body); err != nil {
		c.Status(http.StatusBadRequest)
		return err
	}
	sessionId, err := api.biz.Login(body.Email, body.Password)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return err
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"data": fiber.Map{
			"session_id": sessionId,
		},
	})
}

type SessionBody struct {
	SessionId string `json:"session_id"`
}

func (api *Api) Ping(c *fiber.Ctx) error {
	body := new(SessionBody)
	if err := c.BodyParser(body); err != nil {
		c.Status(http.StatusBadRequest)
		return err
	}
	count, err := api.biz.Ping(body.SessionId)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return err
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"data": fiber.Map{
			"count":   count,
			"message": "Pong",
		},
	})
}

func (api *Api) CountPing(c *fiber.Ctx) error {
	body := new(SessionBody)
	if err := c.BodyParser(body); err != nil {
		c.Status(http.StatusBadRequest)
		return err
	}
	user, err := api.biz.GetUserFromSession(body.SessionId)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return err
	}
	id := int64(user.ID)
	count, err := api.biz.GetPingCount(c.Context(), id)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return err
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"data": fiber.Map{"count": count},
	})
}

func (api *Api) CountUser(c *fiber.Ctx) error {
	count, err := api.biz.GetCountUser()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return err
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"data": fiber.Map{"user": count},
	})
}

func (api *Api) TopUser(c *fiber.Ctx) error {
	limitSrt := c.Params("limit", "10")
	limit, err := strconv.Atoi(limitSrt)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return err
	}
	topUser, err := api.biz.GetTopUser(limit)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return err
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"data": topUser,
	})
}

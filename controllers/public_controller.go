package controllers

import (
	"encoding/json"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"wahyuade.com/simple-job-api/helpers"
	"wahyuade.com/simple-job-api/models"
	"wahyuade.com/simple-job-api/repositories"
)

type Public struct {
	repo repositories.Repositories
}

func (p Public) Register(c *fiber.Ctx) error {
	register_rule := make(map[string][]string)
	register_rule["name"] = []string{"required"}
	register_rule["username"] = []string{"required"}
	register_rule["password"] = []string{"required"}

	request := make(map[string]interface{})
	json.Unmarshal(c.Body(), &request)

	status, errValidation, cleanRequest := helpers.Validate(request, register_rule)
	if !status {
		return c.Status(400).JSON(models.APIResponse{
			Status:  false,
			Message: "Invalid request",
			Error:   errValidation,
		})
	}

	user := models.User{}
	user.Name = cleanRequest["name"].(string)
	user.Username = cleanRequest["username"].(string)
	user.Password = cleanRequest["password"].(string)

	if p.repo.UserRespository.IsUsernameExist(user.Username) {
		err := make(map[string]string)
		err["username"] = "already exists"
		return c.Status(400).JSON(models.APIResponse{
			Status:  false,
			Message: "Failed to register user",
			Error:   err,
		})
	}
	p.repo.UserRespository.Insert(user)
	return c.JSON(models.APIResponse{
		Status:  true,
		Message: "Success",
	})
}

func (p Public) Login(c *fiber.Ctx) error {
	login_rule := make(map[string][]string)
	login_rule["username"] = []string{"required"}
	login_rule["password"] = []string{"required"}

	request := make(map[string]interface{})
	json.Unmarshal(c.Body(), &request)

	status, errValidation, cleanRequest := helpers.Validate(request, login_rule)
	if !status {
		return c.Status(400).JSON(models.APIResponse{
			Status:  false,
			Message: "Invalid request",
			Error:   errValidation,
		})
	}
	user := p.repo.UserRespository.GetByUsername(cleanRequest["username"].(string))
	password := cleanRequest["password"].(string)
	isPasswordMatch, _ := helpers.MatchHash(password, user.Password)
	if !isPasswordMatch {
		return c.Status(400).JSON(models.APIResponse{
			Status:  false,
			Message: "Invalid username and password",
		})
	}
	claims := jwt.MapClaims{
		"id":       user.Id,
		"name":     user.Name,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(models.APIResponse{
		Status:  true,
		Message: "Success",
		Data:    fiber.Map{"token": t},
	})
}

func NewPublic(r repositories.Repositories) Public {
	return Public{
		repo: r,
	}
}

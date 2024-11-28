package server

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"backend/internal/database"
)

type FiberServer struct {
	*fiber.App

	DB       *gorm.DB
	Validate *validator.Validate
}

func New() *FiberServer {

	validator := validator.New(validator.WithPrivateFieldValidation())
	validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})

	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "backend",
			AppName:      "backend",
		}),

		Validate: validator,
		DB:       database.New(),
	}

	return server
}

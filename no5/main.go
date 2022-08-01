package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)

func main()  {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, X-Request-With, Content-Type, Accept, nds-token",
		AllowMethods:     "HEAD, GET, POST, PUT, DELETE, OPTIONS",
		AllowCredentials: true,
	}))

	app.Get("/", healthCheckService)


	err := app.Listen("0.0.0.0:8080")
	if err != nil {
		log.Fatalln("error :" , err.Error())
	}
}

func healthCheckService(c *fiber.Ctx) error{
	return RestSuccess(c, "healthy", 200, nil)
}

type restSuccess struct {
	ErrMessage string      `json:"message"`
	ErrStatus  int         `json:"status"`
	Data       interface{} `json:"data"`
}
func RestSuccess(c *fiber.Ctx, message string, status int, data interface{}) error {
	return c.Status(status).JSON(restSuccess{
		ErrMessage: message,
		ErrStatus:  status,
		Data:       data,
	})
}

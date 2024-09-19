package main

import (
	"encoding/json"
	"errors"
	_ "fiber-httpbin/docs"
	"fiber-httpbin/routes"
	"fiber-httpbin/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"log"
)

// @title Fiber HTTPBin API
// @version 1.0
// @description A web service similar to httpbin.org built with Fiber.
// @host localhost:3000
// @BasePath /
func main() {
	app := fiber.New(fiber.Config{
		JSONDecoder:    json.Unmarshal,
		JSONEncoder:    json.Marshal,
		BodyLimit:      10 * 1024 * 1024,
		ReadBufferSize: 10 * 1024 * 1024,
		// Override default error handler
		ErrorHandler: processError,
	})

	// Configure CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS,PATCH,HEAD",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))
	app.Use(etag.New())
	app.Use(logger.New())

	// Add recovery middleware
	app.Use(recover.New())

	// Swagger route
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Setup routing
	routes.Setup(app)

	// Start server
	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}

func processUtilError(ctx *fiber.Ctx, err *utils.Error) error {
	return ctx.Status(err.Status).JSON(err)
}

func processFiberError(ctx *fiber.Ctx, err *fiber.Error) error {
	return ctx.Status(err.Code).JSON(utils.Error{Status: err.Code, Code: "internal-server", Message: err.Message})
}

func processDefaultError(ctx *fiber.Ctx, err error) error {
	return ctx.Status(500).JSON(utils.Error{Status: 500, Code: "internal-server", Message: err.Error()})
}

func processError(ctx *fiber.Ctx, err error) error {
	var e *utils.Error
	if errors.As(err, &e) {
		return processUtilError(ctx, e)
	}

	var f *fiber.Error
	if errors.As(err, &f) {
		return processFiberError(ctx, f)
	}

	return processDefaultError(ctx, err)
}

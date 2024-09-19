package routes

import (
	"fiber-httpbin/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	// Define routes
	app.Get("/get", controllers.GetHandler)
	app.Post("/post", controllers.PostHandler)
	app.Get("/headers", controllers.HeadersHandler)
	app.Get("/ip", controllers.IPHandler)
	app.Get("/delay", controllers.DelayHandler)
	app.Get("/status", controllers.StatusHandler)
	app.All("/anything", controllers.AnythingHandler)

}

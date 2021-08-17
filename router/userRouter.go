package router

import (
   "github.com/gofiber/fiber/v2"
   "api/service"
)

func UserRoutes(app *fiber.App) {
   // Grouping User Routes
   router := app.Group("/users")
   
   // Routes...
   router.Get("/", service.GetUsers)
   router.Get("/:id", service.GetUser)
   router.Post("/", service.CreateUser)
   router.Put("/:id", service.UpdateUser)
   router.Delete("/:id", service.DeleteUser)
}
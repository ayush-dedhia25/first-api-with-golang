package router

import (
   "github.com/gofiber/fiber/v2"
   "api/service"
)

func UserRoutes(app *fiber.App) {
   // Grouping User Routes
   userRoutes := app.Group("/users")
   
   // Routes...
   userRoutes.Get("/", service.GetUsers)
   userRoutes.Get("/:id", service.GetUser)
   userRoutes.Post("/", service.CreateUser)
   userRoutes.Put("/:id", service.UpdateUser)
   userRoutes.Delete("/:id", service.DeleteUser)
}
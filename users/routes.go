package users

import (
   "github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
   app.Get("/users", GetAllUsers)
   app.Get("/users/:id", GetUser)
   app.Post("/users/:id", CreateUser)
   app.Put("/users/:id", UpdateUser)
   app.Delete("/users/:id", DeleteUser)
}
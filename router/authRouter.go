package router

import (
   "github.com/gofiber/fiber/v2"
   "api/service"
)

func AuthRouter(app *fiber.App) {
   auth := app.Group("/auth")
   
   auth.Get("/login", service.Login)
   auth.Get("/logout", service.Logout)
}
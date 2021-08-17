package router

import (
   "github.com/gofiber/fiber/v2"
   "api/service"
)

func AuthRouter(app *fiber.App) {
   router := app.Group("/auth")
   
   router.Get("/login", service.Login)
   router.Get("/logout", service.Logout)
}
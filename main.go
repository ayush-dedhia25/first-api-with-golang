package main

import (
   "github.com/gofiber/fiber/v2"
   "api/users"
)

func main() {
   // Instantiate Fiber Instance (Main)
   app := fiber.New()
   
   // Routes
   users.UserRoutes(app)
   
   // Serving and Listening the Port and Server
   app.Listen(":4000")
}
package main

import (
   "github.com/gofiber/fiber/v2"
   "github.com/gofiber/fiber/v2/middleware/logger"
   "api/database"
   "api/router"
)

func main() {
   // Main Fiber Entry Part
   app := fiber.New()
   
   // Middlewares
   app.Use(logger.New(logger.Config{
      Next: nil,
      Format: "[${time}] ${status} - ${latency} ${method} ${path}\n",
      TimeFormat: "15:04:05",
   }))
   
   // Initializing Database
   database.InitDatabase()
   
   // Routers Here...
   router.UserRoutes(app)
   
   // Serving and Listening to the Port
   app.Listen(":3000")
}
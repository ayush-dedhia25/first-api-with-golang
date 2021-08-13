package main

import (
   "fmt"
   
   "github.com/gofiber/fiber/v2"
   "gorm.io/gorm"
   "gorm.io/driver/sqlite"
   
   "api/db"
   "api/router"
   "api/model"
)

func initDatabase() {
   var err error
   db.DBConn, err = gorm.Open(sqlite.Open("./store.db"), &gorm.Config{})
   if err != nil {
      panic("Failed to connect to database!")
   }
   db.DBConn.AutoMigrate(&model.User{})
   fmt.Println("Database connection established!")
}

func main() {
   // Main Fiber Entry Part
   app := fiber.New()
   
   // Initializing Database
   initDatabase()
   
   // Routers Here...
   router.UserRoutes(app)
   
   // Serving and Listening to the Port
   app.Listen(":3000")
}
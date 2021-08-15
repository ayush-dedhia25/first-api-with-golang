package database

import (
   "fmt"
   "gorm.io/driver/sqlite"
   "gorm.io/gorm"
   "api/model"
)

var (
   DBConn *gorm.DB
)

func InitDatabase() {
   var err error
   DBConn, err = gorm.Open(sqlite.Open("./store.db"), &gorm.Config{})
   if err != nil {
      panic("Failed to connect to database!")
   }
   DBConn.AutoMigrate(&model.User{})
   fmt.Println("Database connection established!")
}
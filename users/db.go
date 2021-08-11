package users

import (
   "gorm.io/gorm"
   "gorm.io/driver/sqlite"
)

func Connect(dbName string) *gorm.DB {
   db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
   if err != nil {
      panic("Failed to connect to database!")
   }
   db.AutoMigrate(&User{})
   return db
}
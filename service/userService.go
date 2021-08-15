package service

import (
   // "fmt"
   "github.com/gofiber/fiber/v2"
   db "api/database"
   "api/model"
)

func GetUsers(ctx *fiber.Ctx) error {
   var users []model.User
   // Retrieving all the users from the database
   db.DBConn.Find(&users)
   
   // Return the success response with the list of users;
   return ctx.Status(200).JSON(fiber.Map{
      "success": true,
      "data": users,
   })
}

func GetUser(ctx *fiber.Ctx) error {
   var user model.User
   // Getting userId from the URL params
   userId := ctx.Params("id")
   
   // Finding the first instance of the user from the database
   db.DBConn.Find(&user, userId)
   
   // Checking if user exists in the database
   if user.Email == "" {
      return ctx.Status(404).SendString("User not available!")
   }
   
   // Returning the success response with the user data;
   return ctx.Status(200).JSON(fiber.Map{
      "success": true,
      "data": user,
   })
}

func CreateUser(ctx *fiber.Ctx) error {
   // Creating new user object
   user := new(model.User)
   
   // Filling values coming from request body into above instance
   // Checking for any error
   if err := ctx.BodyParser(&user); err != nil {
      return ctx.Status(500).SendString(err.Error())
   }
   
   // Selecting the user from the database to see
   // if is exists in the database
   db.DBConn.Take(&user)
   
   // Saving new user into the database if the user doesn't exists!
   if user.Email == "" {
      db.DBConn.Create(&user)
      
      // Returning the response
      return ctx.Status(200).JSON(fiber.Map{
         "success": true,
         "message": "User Created!",
      })
   }
   
   return ctx.Status(409).SendString("User already exists!")
}

func UpdateUser(ctx *fiber.Ctx) error {
   // Getting userId from the URL params
   userId := ctx.Params("id")
   
   // Creating new user object
   user := new(model.User)
   
   // Finding the user in the database
   db.DBConn.First(&user, userId)
   
   // Checking if user exists?
   if user.Email == "" {
      return ctx.Status(404).SendString("User not found!")
   }
   
   // Filling values coming from request body into above user instance
   // Checking for any error!?
   if err := ctx.BodyParser(&user); err != nil {
      return ctx.Status(500).SendString(err.Error())
   }
   
   // Updating the user
   db.DBConn.Save(&user)
   
   // Returning the success response to the end user;
   return ctx.Status(200).JSON(fiber.Map{
      "success": true,
      "message": "User updated!",
   })
}

func DeleteUser(ctx *fiber.Ctx) error {
   var user model.User
   // Getting userId from the URL params
   userId := ctx.Params("id")
   
   // Find the first instance of user from the database
   db.DBConn.First(&user, userId)
   
   // Checking if user exists!?
   if user.Email == "" {
      return ctx.Status(404).SendString("User not found!")
   }
   
   // Deleting the particular user from the database
   db.DBConn.Delete(&user)
   
   // Returning the success response to the end user;
   return ctx.Status(204).JSON(fiber.Map{
      "success": true,
      "message": "User deleted!",
   })
}
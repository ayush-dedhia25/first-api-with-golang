package users

import (
   // "fmt"
   "github.com/gofiber/fiber/v2"
)

var db = Connect("../test.db")

func GetAllUsers(c *fiber.Ctx) error {
   var users []User
   // Retrieving all the users from the database
   db.Find(&users)
   
   // Return the success response with the list of users;
   return c.Status(200).JSON(fiber.Map{
      "success": true,
      "data": users,
   })
}

func GetUser(c *fiber.Ctx) error {
   var user User
   // Getting userId from the URL params
   userId := c.Params("id")
   
   // Finding the first instance of the user from the database
   db.Find(&user, userId)
   
   // Checking if user exists in the database
   if user.Email == "" {
      return c.Status(404).SendString("User not available!")
   }
   
   // Returning the success response with the user data;
   return c.Status(200).JSON(fiber.Map{
      "success": true,
      "data": user,
   })
}

func CreateUser(c *fiber.Ctx) error {
   // Creating new user object
   user := new(User)
   
   // Filling values coming from request body into above instance
   // Checking for any error
   if err := c.BodyParser(&user); err != nil {
      return c.Status(500).SendString(err.Error())
   }
   
   // Saving new user into the database
   db.Create(&user)
   
   // Returning the response
   return c.Status(200).JSON(fiber.Map{
      "success": true,
      "message": "User Created!",
   })
}

func UpdateUser(c *fiber.Ctx) error {
   // Getting userId from the URL params
   userId := c.Params("id")
   
   // Creating new user object
   user := new(User)
   
   // Finding the user in the database
   db.First(&user, userId)
   
   // Checking if user exists?
   if user.Email == "" {
      return c.Status(404).SendString("User not found!")
   }
   
   // Filling values coming from request body into above user instance
   // Checking for any error!?
   if err := c.BodyParser(&user); err != nil {
      return c.Status(500).SendString(err.Error())
   }
   
   // Updating the user
   db.Save(&user)
   
   // Returning the success response to the end user;
   return c.Status(200).JSON(fiber.Map{
      "success": true,
      "message": "User updated!",
   })
}

func DeleteUser(c *fiber.Ctx) error {
   var user User
   // Getting userId from the URL params
   userId := c.Params("id")
   
   // Find the first instance of user from the database
   db.First(&user, userId)
   
   // Checking if user exists!?
   if user.Email == "" {
      return c.Status(404).SendString("User not found!")
   }
   
   // Deleting the particular user from the database
   db.Delete(&user)
   
   // Returning the success response to the end user;
   return c.Status(200).JSON(fiber.Map{
      "success": true,
      "message": "User deleted!",
   })
}
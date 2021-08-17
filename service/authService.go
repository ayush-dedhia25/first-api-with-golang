package service

import (
   "time"
   "github.com/gofiber/fiber/v2"
   "github.com/dgrijalva/jwt-go"
   db "api/database"
   "api/model"
)

var jwtTokenSecret = []byte("reallySecretToken")

type Payload struct {
   Email    string
   Password string
   jwt.StandardClaims
}

func Login(ctx *fiber.Ctx) error {
   var credentials model.User
   if err := ctx.BodyParser(&credentials); err != nil {
      return ctx.Status(500).SendString("Error!")
   }
   
   db.DBConn.First(&credentials)
   if credentials.Email == "" {
      return ctx.Status(403).SendString("You are not a member!")
   }
   
   expirationTime := time.Now().Add(time.Minute * 5)
   
   payload := &Payload{
      Password: credentials.Name,
      Email: credentials.Email,
      StandardClaims: jwt.StandardClaims{
         ExpiresAt: expirationTime.Unix(),
      },
   }
   
   token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
   tokenString, err := token.SignedString(jwtTokenSecret)
   if err != nil {
      return ctx.Status(500).SendString("Error with signing token!")
   }
   
   userCookie := new(fiber.Cookie)
   userCookie.Name = "token"
   userCookie.Value = tokenString
   userCookie.Expires = expirationTime
   ctx.Cookie(userCookie)
   
   return ctx.Status(200).SendString("Login!!!")
}

func Logout(ctx *fiber.Ctx) error {
   return ctx.Status(204).JSON(fiber.Map{})
}
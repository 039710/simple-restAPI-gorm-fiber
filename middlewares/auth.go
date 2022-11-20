// create middleware for jwt auth

package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func Auth() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// get token from header
		token := c.Get("Authorization")
		// remove Bearer from token
		token = token[7:]
		// validate token
		_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, nil
			}
			return []byte("secret"), nil
		})
		// if token is invalid
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}
		// if token is valid
		return c.Next()
	}
}

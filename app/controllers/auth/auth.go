package auth

import (
	"learn-go-fiber/app/helpers"
	"learn-go-fiber/app/models"
	"learn-go-fiber/app/requests"
	"learn-go-fiber/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
	"golang.org/x/crypto/bcrypt"
)

var DB = config.DB

func LoginUser(c *fiber.Ctx) error {
	p := new(requests.LoginForm)
	if err := c.BodyParser(p); err != nil {
		return err
	}
	v := validate.Struct(p)
	if !v.Validate() {
		return c.JSON(fiber.Map{
			"message": v.Errors.One(),
		})
	}
	u := new(models.User)

	if res := DB.Where("email = ?", p.Email).First(&u); res.RowsAffected <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid Email!",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p.Password)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Password is incorrect!",
		})
	}

	accessToken, refreshToken := helpers.GenerateTokens(u.UUID.String())
	accessCookie, refreshCookie := helpers.GetAuthCookies(accessToken, refreshToken)

	c.Cookie(accessCookie)
	c.Cookie(refreshCookie)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})

}

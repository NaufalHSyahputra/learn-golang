package auth

import (
	"os"
	"time"

	"github.com/form3tech-oss/jwt-go"
)

// func Login(c *fiber.Ctx) error {
// 	loginRequest := new(requests.LoginForm)
// 	if err := c.BodyParser(loginRequest); err != nil {
// 		return err
// 	}
// 	v := validate.Struct(loginRequest)
// 	if !v.Validate() {
// 		return c.JSON(fiber.Map{
// 			"message": v.Errors.One(),
// 		})
// 	}
// 	db := config.GetDBInstance()
// 	u := new(models.User)
// 	if db.Model(&models.User{}).Where("email = ?", u.Email).First(&u).RowsAffected <= 0 {

// 	}
// }

func CreateToken(userid uint64) (string, error) {
	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

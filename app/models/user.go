package models

import (
	"learn-go-fiber/config"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Base
	Email    string `json:"email"`
	Password string `json:"password" gorm:"->:false;<-:create"`
	Name     string `json:"name"`
	Todos    []Todo
}

// Claims represent the structure of the JWT token
type Claims struct {
	jwt.StandardClaims
	ID uint `gorm:"primaryKey"`
}

// HashPassword encrypts user password
func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	user.Password = string(bytes)

	return nil
}

// CheckPassword checks user password
func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}

	return nil
}

// CreateUserRecord creates a user record in the database
func (user *User) CreateUser() error {
	result := config.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

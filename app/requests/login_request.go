package requests

import "github.com/gookit/validate"

type LoginForm struct {
	Email    string `json:"email" xml:"email" form:"email" validate:"required|email"`
	Password string `json:"password" xml:"password" form:"password" validate:"required"`
}

// Messages you can custom validator error messages.
func (f LoginForm) Messages() map[string]string {
	return validate.MS{
		"required":    "{field} wajib diisi.",
		"Email.email": "Format {field} tidak valid.",
	}
}

// Translates you can custom field translates.
func (f LoginForm) Translates() map[string]string {
	return validate.MS{
		"Email":    "Email",
		"Password": "Password",
	}
}

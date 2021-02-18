package requests

import "github.com/gookit/validate"

type TodoForm struct {
	Name        string `json:"name" xml:"name" form:"name" validate:"required|minLen:4"`
	IsCompleted bool   `json:"is_complete" xml:"is_complete" form:"is_complete" validate:"required|isBool"`
}

// Messages you can custom validator error messages.
func (f TodoForm) Messages() map[string]string {
	return validate.MS{
		"required":           "{field} wajib diisi.",
		"Name.minLen":        "{field} harus lebih dari 4 karakter",
		"IsCompleted.isBool": "{field} hanya boleh diisi true atau false",
	}
}

// Translates you can custom field translates.
func (f TodoForm) Translates() map[string]string {
	return validate.MS{
		"Name":        "Deskripsi Todo",
		"IsCompleted": "Status Todo",
	}
}

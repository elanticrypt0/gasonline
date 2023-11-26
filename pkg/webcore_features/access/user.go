package access

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID   string
	Name   string `json:"name"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
	Status uint8  `json:"status"`
}

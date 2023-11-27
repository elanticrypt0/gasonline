package access

import "gorm.io/gorm"

// Authentication factors

type AccessFactor struct {
	gorm.Model
	Email       bool
	Phone       bool
	Token       bool
	CodeTemp    string
	TokenActive string
}

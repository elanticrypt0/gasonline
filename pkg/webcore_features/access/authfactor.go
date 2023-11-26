package access

import "gorm.io/gorm"

// Authentication factors

type AuthFactor struct {
	gorm.Model
	Email       bool
	Phone       bool
	Token       bool
	CodeTemp    string
	TokenActive string
}

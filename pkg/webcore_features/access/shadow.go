package access

import "gorm.io/gorm"

type Shadow struct {
	gorm.Model
	// Salt   string // the salt
	Shadow string // encrypted password
}

package access

import "gorm.io/gorm"

type GroupPerms struct {
	gorm.Model
	GroupID uint
	Group   Group
	Write   bool   `json:"write"`
	Read    bool   `json:"read"`
	Path    string `json:"path"`
}

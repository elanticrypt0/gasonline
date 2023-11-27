package access

import (
	"log"

	"github.com/elanticrypt0/go4it"
	"gorm.io/gorm"
)

const seedDir = "./seeds/"

func SeedGroups(db *gorm.DB) {
	group_list := []Group{}
	go4it.ReadAndParseJson(seedDir+"groups", &group_list)
	db.Save(&group_list)
	log.Println("Groups seeded")
}

func SeedGroupsPerms(db *gorm.DB) {
	perms_list := []GroupPerms{}
	go4it.ReadAndParseJson(seedDir+"groupperms", &perms_list)
	db.Save(&perms_list)
	log.Println("Groups Perms seeded")
}

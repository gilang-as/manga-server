package config

import (
	"github.com/jinzhu/gorm"
	"manga-server/domain/models"
)

func DBMigrate(db *gorm.DB) {
	db.AutoMigrate(models.Manga{}).Exec("ALTER TABLE `manga` CHANGE `original_title` `original_title` VARCHAR(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL").Exec("ALTER TABLE `manga` CHANGE `synopsis` `synopsis` TEXT CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL")

}

func DropAll(db *gorm.DB) {
	db.DropTableIfExists(models.Manga{})
}
package migration

import (
	"vinbigdata/internal/repository/entity"

	"gorm.io/gorm"
)

func CreateTable(db *gorm.DB) {
	db.AutoMigrate(&entity.UserCalls{})
}

func DropTable(db *gorm.DB) {
	db.Migrator().DropTable(&entity.UserCalls{})

}

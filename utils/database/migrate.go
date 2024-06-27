package database

import (
	DataUser "backend-api/features/users/data"
	"fmt"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(DataUser.User{})
	fmt.Println("\n[MIGRATION] Success creating user data")

}

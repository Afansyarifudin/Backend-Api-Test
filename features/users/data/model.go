package data

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	Name        string `gorm:"column:name;type:varchar(255)"`
	Email       string `gorm:"column:email;unique;type:varchar(255)"`
	Password    string `gorm:"column:password;type:varchar(255)"`
	DateOfBirth string `gorm:"column:date_of_birth"`
	Age         string `gorm:"column:age"`
}

func (User) TabeleName() string {
	return "users"
}

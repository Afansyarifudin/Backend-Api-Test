package data

import (
	"backend-api/features/users"
	"errors"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserData struct {
	db *gorm.DB
}

func New(db *gorm.DB) users.UserDataInterface {
	return &UserData{
		db: db,
	}
}

func (ud *UserData) Register(newData users.User) (*users.User, error) {
	var dbData = new(User)
	dbData.Name = newData.Name
	dbData.Email = newData.Email
	dbData.Password = newData.Password
	dbData.DateOfBirth = newData.DateOfBirth
	dbData.Age = newData.Age

	if err := ud.db.Create(dbData).Error; err != nil {
		return nil, err
	}

	newData.ID = dbData.ID

	return &newData, nil
}

func (ud *UserData) Login(email, password string) (*users.User, error) {
	var dbData = new(User)
	dbData.Email = email

	var qry = ud.db.Where("email ?", dbData.Email).
		First(dbData)

	var dataCount int64
	qry.Count(&dataCount)
	if dataCount == 0 {
		return nil, errors.New("credentials not found")
	}

	if err := qry.Error; err != nil {
		logrus.Info("D: Error: ", err.Error())
		return nil, err
	}

	passwordByted := []byte(password)

	err := bcrypt.CompareHashAndPassword([]byte(dbData.Password), []byte(passwordByted))
	if err != nil {
		logrus.Info("Incorrect Password")
		return nil, errors.New("incorrect password")
	}

	var result = new(users.User)
	result.ID = dbData.ID
	result.Email = dbData.Email
	result.Name = dbData.Name
	result.DateOfBirth = dbData.DateOfBirth
	result.Age = dbData.Age

	return result, nil

}

func (ud *UserData) CreateUser(newData users.User) (*users.User, error) {
	var dbData = new(User)
	dbData.Name = newData.Name
	dbData.Email = newData.Email
	dbData.Password = newData.Password
	dbData.DateOfBirth = newData.DateOfBirth
	dbData.Age = newData.Age

	if err := ud.db.Create(dbData).Error; err != nil {
		return nil, err
	}

	newData.ID = dbData.ID

	return &newData, nil
}

func (ud *UserData) GetAllUsers() ([]users.User, error) {
	var users []users.User

	qry := ud.db.Table("users").
		Where("users.deleted_at is null")

	if err := qry.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (ud *UserData) GetUserByID(id uint) (*users.User, error) {
	var user users.User
	var qry = ud.db.Table("users").Select("users.*").
		Where("users.id = ?", id).
		Where("users.deleted_at is null").
		Scan(&user)

	if err := qry.Error; err != nil {
		return &user, err
	}

	return &user, nil
}

func (ud *UserData) UpdateUser(newData users.UpdateUser, id uint) (bool, error) {
	var qry = ud.db.Table("users").
		Where("id = ?", id).Updates(User{
		Name:        newData.Name,
		Email:       newData.Email,
		DateOfBirth: newData.DateOfBirth,
		Age:         newData.Age,
	})

	if err := qry.Error; err != nil {
		return false, err
	}

	if dataCount := qry.RowsAffected; dataCount < 1 {
		return false, errors.New("update data error")
	}

	return true, nil

}

func (ud *UserData) DeleteUser(id uint) (bool, error) {
	var deleteData users.User

	if err := ud.db.Table("users").Where("id = ?", id).Delete(deleteData).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (ud *UserData) GetByEmail(email string) (*users.User, error) {
	var dbData = new(User)
	dbData.Email = email

	if err := ud.db.Where("email = ?", dbData.Email).First(dbData).Error; err != nil {
		logrus.Info("DB Error : ", err.Error())
		return nil, err
	}

	var result = new(users.User)
	result.ID = dbData.ID
	result.Name = dbData.Name
	result.Email = dbData.Email
	result.DateOfBirth = dbData.DateOfBirth
	result.Age = dbData.Age

	return result, nil

}

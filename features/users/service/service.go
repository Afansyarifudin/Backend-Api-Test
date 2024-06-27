package service

import (
	"backend-api/features/users"
	"backend-api/helper"
	"backend-api/helper/encrypt"
	"errors"
	"strings"
)

type UserSevice struct {
	d users.UserDataInterface
	e encrypt.HashInterface
	j helper.JWTInterface
}

func New(data users.UserDataInterface, encrypt encrypt.HashInterface, jwt helper.JWTInterface) users.UserServiceInterface {
	return &UserSevice{
		d: data,
		e: encrypt,
		j: jwt,
	}
}

func (us *UserSevice) Register(newData users.User) (*users.User, error) {
	_, err := us.d.GetByEmail(newData.Email)
	if err == nil {
		return nil, errors.New("email already registered")
	}
	hashPassword, err := us.e.HashPassword(newData.Password)
	if newData.Password != "" {
		if err != nil {
			return nil, errors.New("error hashing password")
		}
	}
	newData.Password = hashPassword
	result, err := us.d.CreateUser(newData)
	if err != nil {
		return nil, errors.New("failed to register user")
	}

	return result, nil
}

func (us *UserSevice) Login(email, password string) (*users.UserCredential, error) {
	result, err := us.d.Login(email, password)
	if err != nil {
		if strings.Contains(err.Error(), "Incorrect Password") {
			return nil, errors.New("incorrect password")
		}
		if strings.Contains(err.Error(), "Not Found") {
			return nil, errors.New("user not found")
		}
		return nil, errors.New("user not found")
	}

	tokenData := us.j.GenerateJWT(result.ID, result.Email)
	if tokenData == nil {
		return nil, errors.New("token process failed")
	}

	response := new(users.UserCredential)
	response.Name = result.Name
	response.Email = result.Email
	response.Access = tokenData

	return response, nil
}

func (us *UserSevice) CreateUser(newData users.User) (*users.User, error) {
	hashPassword, err := us.e.HashPassword(newData.Password)
	if newData.Password != "" {
		if err != nil {
			return nil, errors.New("error hashing password")
		}
	}
	newData.Password = hashPassword
	result, err := us.d.CreateUser(newData)
	if err != nil {
		return nil, errors.New("failed to create user")
	}

	return result, nil

}

func (us *UserSevice) GetAllUsers() ([]users.User, error) {
	result, err := us.d.GetAllUsers()
	if err != nil {
		return nil, errors.New("get all users failed")
	}
	return result, nil
}

func (us *UserSevice) GetUserByID(id uint) (*users.User, error) {
	result, err := us.d.GetUserByID(id)
	if err != nil {
		return nil, errors.New("get user by id failed")
	}
	return result, nil
}

func (us *UserSevice) UpdateUser(newData users.UpdateUser, id uint) (bool, error) {
	hashPassword, err := us.e.HashPassword(newData.Password)
	if newData.Password != "" {
		if err != nil {
			return false, errors.New("error hashing password")
		}
	}
	newData.Password = hashPassword
	result, err := us.d.UpdateUser(newData, id)
	if err != nil {
		return false, errors.New("update user failed")
	}
	return result, nil
}

func (us *UserSevice) DeleteUser(id uint) (bool, error) {
	result, err := us.d.DeleteUser(id)
	if err != nil {
		return false, errors.New("delete user failed")
	}
	return result, nil
}

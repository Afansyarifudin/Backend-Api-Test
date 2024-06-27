package users

import (
	"github.com/labstack/echo/v4"
)

type User struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	DateOfBirth string `json:"date_of_birth"`
	Age         string `json:"age"`
}

type UserCredential struct {
	Name   string         `json:"name"`
	Email  string         `json:"email"`
	Access map[string]any `json:"token"`
}

type UpdateUser struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	DateOfBirth string `json:"date_of_birth"`
	Age         string `json:"age"`
}

type UserHandlerInterface interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	GetAllUsers() echo.HandlerFunc
	GetUserByID() echo.HandlerFunc
	UpdateUser() echo.HandlerFunc
	CreateUser() echo.HandlerFunc
	DeleteUser() echo.HandlerFunc
}

type UserServiceInterface interface {
	Register(newData User) (*User, error)
	Login(email, password string) (*UserCredential, error)
	CreateUser(newData User) (*User, error)
	GetAllUsers() ([]User, error)
	GetUserByID(id uint) (*User, error)
	UpdateUser(newData UpdateUser, id uint) (bool, error)
	DeleteUser(id uint) (bool, error)
}

type UserDataInterface interface {
	Register(newData User) (*User, error)
	Login(email, password string) (*User, error)
	GetByEmail(email string) (*User, error)
	CreateUser(newData User) (*User, error)
	GetAllUsers() ([]User, error)
	GetUserByID(id uint) (*User, error)
	UpdateUser(newData UpdateUser, id uint) (bool, error)
	DeleteUser(id uint) (bool, error)
}

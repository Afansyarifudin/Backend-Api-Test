package handler

type UserResponse struct {
	ID          uint   `json:"id" form:"id"`
	Name        string `json:"name" form:"name"`
	Email       string `json:"email" form:"email"`
	Age         string `json:"age"`
	DateOfBirth string `json:"date_of_birth"`
}

type LoginResponse struct {
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Token any    `json:"token"`
}
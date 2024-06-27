package handler

type UpdateUserRequest struct {
	Name        string `json:"name" form:"name"`
	Email       string `json:"email" form:"email"`
	Password    string `json:"password" form:"password"`
	DateOfBirth string `json:"date_of_birth" form:"date_of_birth"`
	Age         string `json:"age" form:"age"`
}

type RegisterInput struct {
	Name        string `json:"name" form:"name" validate:"required"`
	Email       string `json:"email" form:"email" validate:"required,email"`
	Password    string `json:"password" form:"password" validate:"required"`
	DateOfBirth string `json:"date_of_birth" form:"date_of_birth" validate:"required"`
	Age         string `json:"age" form:"age" validate:"required"`
}

type LoginInput struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

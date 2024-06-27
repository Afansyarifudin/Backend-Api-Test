package handler

import (
	"backend-api/features/users"
	"backend-api/helper"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	s   users.UserServiceInterface
	jwt helper.JWTInterface
}

func NewHandler(service users.UserServiceInterface, jwt helper.JWTInterface) users.UserHandlerInterface {
	return &UserHandler{
		s:   service,
		jwt: jwt,
	}
}

// Register creates a new user account.
// @Summary Register
// @Description Register a new user with name, email, date of birth, age, and password
// @Tags Authentication
// @Accept json
// @Produce json
// @Param input body RegisterInput true "Register input"
// @Success 201 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Failure 500 {object} helper.Response
// @Router /register [post]
func (uh *UserHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(RegisterInput)
		if err := c.Bind(&input); err != nil {
			c.Logger().Info("Handler: Bind input error: ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, "error bind input", nil))
		}

		isValid, errors := helper.ValidateForm(input)
		if !isValid {
			return c.JSON(http.StatusBadRequest, helper.FormatResponseValidation(false, "error validation", errors))
		}

		var userRequest = new(users.User)
		userRequest.Name = input.Name
		userRequest.Email = input.Email
		userRequest.DateOfBirth = input.DateOfBirth
		userRequest.Age = input.Age
		userRequest.Password = input.Password

		result, err := uh.s.Register(*userRequest)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, err.Error(), nil))
		}

		var response = new(UserResponse)
		response.ID = result.ID
		response.Name = result.Name
		response.Email = result.Email
		response.Age = result.Age
		response.DateOfBirth = result.DateOfBirth
		return c.JSON(http.StatusCreated, helper.FormatResponse(true, "success register user", response))
	}
}

// Login logs a user in.
// @Summary Login
// @Description Login with email and password
// @Tags Authentication
// @Accept json
// @Produce json
// @Param input body LoginInput true "Login input"
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Failure 404 {object} helper.Response
// @Router /login [post]
func (uh *UserHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(LoginInput)

		if err := c.Bind(input); err != nil {
			c.Logger().Info("Handler: Bind input error: ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, "error bind input", nil))
		}

		result, err := uh.s.Login(input.Email, input.Password)

		if err != nil {
			c.Logger().Info("Handler: Login process error: ", err.Error())
			if strings.Contains(err.Error(), "Not Found") {
				return c.JSON(http.StatusNotFound, helper.FormatResponse(false, "credentials not found", nil))
			}
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, err.Error(), nil))
		}

		var response = new(LoginResponse)
		response.Name = result.Name
		response.Email = result.Email
		response.Token = result.Access

		return c.JSON(http.StatusOK, helper.FormatResponse(true, "succes login", response))
	}
}

// GetAllUsers retrieves all users.
// @Summary Get All Users
// @Description Get a list of all users
// @Tags Users
// @Produce json
// @Success 200 {object} helper.Response
// @Failure 500 {object} helper.Response
// @Router /admin/users [get]
func (uh *UserHandler) GetAllUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		result, err := uh.s.GetAllUsers()
		if err != nil {
			c.Logger().Error("Handler : Error to get product : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, err.Error(), nil))
		}

		if len(result) == 0 {
			return c.JSON(http.StatusOK, helper.FormatResponse(true, "user data is empty", nil))
		}

		var userResponseList []UserResponse
		for _, user := range result {
			userResponseList = append(userResponseList,
				UserResponse{
					ID:          user.ID,
					Name:        user.Name,
					Email:       user.Email,
					Age:         user.Age,
					DateOfBirth: user.DateOfBirth,
				})
		}

		return c.JSON(http.StatusOK, helper.FormatResponse(true, "succes get data", userResponseList))
	}
}

// GetUserByID retrieves a user by ID.
// @Summary Get User By ID
// @Description Get a user by their ID
// @Tags Users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Failure 404 {object} helper.Response
// @Router /admin/users/{id} [get]
func (uh *UserHandler) GetUserByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramID = c.Param("id")

		id, err := strconv.Atoi(paramID)
		if err != nil {
			c.Logger().Error("Handler : Param ID error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, "paramId error", nil))
		}

		result, err := uh.s.GetUserByID(uint(id))
		if err != nil {
			return c.JSON(http.StatusNotFound, helper.FormatResponse(false, "get user by id failed", nil))
		}
		var response = new(UserResponse)
		response.ID = result.ID
		response.Name = result.Name
		response.Email = result.Email
		response.Age = result.Age
		response.DateOfBirth = result.DateOfBirth

		return c.JSON(http.StatusOK, helper.FormatResponse(true, "success get user data by id", response))

	}
}

// UpdateUser updates an existing user.
// @Summary Update User
// @Description Update an existing user with name, email, date of birth, age, and password
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param input body UpdateUserRequest true "Update user input"
// @Success 200 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Failure 500 {object} helper.Response
// @Router /admin/users/{id} [put]
func (uh *UserHandler) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(UpdateUserRequest)
		var paramID = c.Param("id")

		id, err := strconv.Atoi(paramID)
		if err != nil {
			c.Logger().Error("Handler : Param ID error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, "paramId error", nil))
		}

		if err := c.Bind(input); err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, "error bind input", nil))

		}

		var userUpdate = new(users.UpdateUser)
		userUpdate.Name = input.Name
		userUpdate.Email = input.Email
		userUpdate.Password = input.Password
		userUpdate.DateOfBirth = input.DateOfBirth
		userUpdate.Age = input.Age

		_, err = uh.s.UpdateUser(*userUpdate, uint(id))

		if err != nil {
			c.Logger().Info("Handler : Input Process Error : ", err.Error())
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, err.Error(), nil))
		}

		return c.JSON(http.StatusOK, helper.FormatResponse(true, "success update data", nil))
	}
}

// CreateUser creates a new user.
// @Summary Create User
// @Description Create a new user with name, email, date of birth, age, and password
// @Tags Users
// @Accept json
// @Produce json
// @Param input body RegisterInput true "Create user input"
// @Success 201 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Failure 500 {object} helper.Response
// @Router /admin/users [post]
func (uh *UserHandler) CreateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(RegisterInput)
		if err := c.Bind(&input); err != nil {
			c.Logger().Info("Handler: Bind input error: ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, "error bind input", nil))
		}

		isValid, errors := helper.ValidateForm(input)
		if !isValid {
			return c.JSON(http.StatusBadRequest, helper.FormatResponseValidation(false, "error validation", errors))
		}

		var userRequest = new(users.User)
		userRequest.Name = input.Name
		userRequest.Email = input.Email
		userRequest.DateOfBirth = input.DateOfBirth
		userRequest.Age = input.Age
		userRequest.Password = input.Password

		result, err := uh.s.CreateUser(*userRequest)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.FormatResponse(false, err.Error(), nil))
		}

		var response = new(UserResponse)
		response.ID = result.ID
		response.Name = result.Name
		response.Email = result.Email
		response.Age = result.Age
		response.DateOfBirth = result.DateOfBirth
		return c.JSON(http.StatusCreated, helper.FormatResponse(true, "success create user", response))

	}
}

// DeleteUser deletes a user by ID.
// @Summary Delete User
// @Description Delete a user by their ID
// @Tags Users
// @Produce json
// @Param id path int true "User ID"
// @Success 201 {object} helper.Response
// @Failure 400 {object} helper.Response
// @Failure 404 {object} helper.Response
// @Router /admin/users/{id} [delete]
func (uh *UserHandler) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var paramID = c.Param("id")

		id, err := strconv.Atoi(paramID)

		if err != nil {
			c.Logger().Error("Handler : Param ID error : ", err.Error())
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, "param id error", nil))
		}

		_, err = uh.s.DeleteUser(uint(id))
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, "error data not found", nil))
		}

		return c.JSON(http.StatusCreated, helper.FormatResponse(true, "success delete data", nil))
	}
}

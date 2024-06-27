package routes

import (
	"backend-api/config"
	"backend-api/features/users"

	"github.com/labstack/echo/v4"
)

// User route
func RouteUser(e *echo.Group, uh users.UserHandlerInterface, cfg config.ProgrammingConfig) {
	// Auth
	e.POST("/register", uh.Register())
	e.POST("/login", uh.Login())

	// Users
	e.POST("/admin/users", uh.CreateUser())
	e.GET("/admin/users", uh.GetAllUsers())
	e.GET("/admin/users/:id", uh.GetUserByID())
	e.PUT("/admin/users/:id", uh.UpdateUser())
	e.DELETE("/admin/users/:id", uh.DeleteUser())
}

// Other route

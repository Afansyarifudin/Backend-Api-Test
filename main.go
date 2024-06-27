package main

import (
	"backend-api/config"
	"backend-api/helper"
	"backend-api/helper/encrypt"
	"backend-api/routes"
	"backend-api/utils/database"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	dataUser "backend-api/features/users/data"
	handlerUser "backend-api/features/users/handler"
	serviceUser "backend-api/features/users/service"

	_ "backend-api/docs"

	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title			REST API Documentation
// @version		1.0
// @description	This is a sample REST_API .
// @termsOfService	http://swagger.io/terms/
// @contact.name	API Support
// @contact.url	http://www.swagger.io/support
// @contact.email	support@swagger.io
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host			localhost:3000
// @BasePath		/api/v1
func main() {
	e := echo.New()
	config := config.InitConfig()

	db, err := database.InitDB(*config)
	if err != nil {
		e.Logger.Fatal("Cannot run database: ", err.Error())
	}

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, "Endpoint not found", nil))
	})

	e.GET("/api", func(c echo.Context) error {
		return c.JSON(http.StatusBadRequest, helper.FormatResponse(false, "Endpoint not found", nil))
	})

	var encrypt = encrypt.New()

	jwtInterface := helper.New(config.Secret, config.RefSecret)

	userModel := dataUser.New(db)

	userServices := serviceUser.New(userModel, encrypt, jwtInterface)
	userController := handlerUser.NewHandler(userServices, jwtInterface)

	group := e.Group("/api/v1")
	routes.RouteUser(group, userController, *config)

	e.GET("/api/v1/docs/swagger/*", echoSwagger.WrapHandler)

	// Automigrate
	database.Migrate(db)

	e.Use(middleware.CORS())

	e.Logger.Debug(db)
	e.Logger.Info(fmt.Sprintf("Listening in port: %d", config.AppPort))
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.AppPort)).Error())

}

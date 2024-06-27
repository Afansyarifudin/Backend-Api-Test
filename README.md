# Backend REST-API Golang

## Description

This is a project of RESTFull API using Go language and detailed tech stack below
| Tech | Used as |
| ------------- | ------------- |
| Echo | web framework in Go |
| Postgresql | Database |
| Swagger | API Spec Docs |
| ORM Gorm | ORM Library in Go |
| Json Web Token (JWT) | encrypt auth |
| Docker | Container Package |

## Library

```go
go get github.com/labstack/echo/v4
go get -u gorm.io/gorm
go get github.com/joho/godotenv
go get github.com/sirupsen/logrus
go get github.com/golang-jwt/jwt/v5
go get github.com/labstack/echo/v4/middleware
go get github.com/labstack/echo-jwt/v4
go get github.com/go-playground/validator/v10
go get github.com/cosmtrek/air
go get -u github.com/swaggo/echo-swagger
go get -d github.com/swaggo/swag/cmd/swag
go install github.com/swaggo/swag/cmd/swag@latest
```

## API List

    http://localhost:3000/api/v1

    Auth Routes :
        POST("/register") register user
        POST("/login") login

    User Routes:
        GET("/admin/users") get all users
        GET("/admin/users/:id") get user by Id
        POST("/admin/users") Create user
        PUT("/admin/users/:id") update user
        DELETE("/admin/users/:id") delete user

## Api docs Swagger

    GET("/docs/swagger/index.html#/") Swagger Docs

    http://localhost:3000/api/v1/docs/swagger/index.html#/

## Mock Data

```Get all users
{
  "status": true,
  "message": "succes get data",
  "data": [
    {
      "id": 1,
      "name": "Nur Afan Syarifudin",
      "email": "afan@gmail.com",
      "age": "4",
      "date_of_birth": "10-20-2021"
    },
    {
      "id": 3,
      "name": "Nur Afan Syarifudin",
      "email": "afan101@gmail.com",
      "age": "4",
      "date_of_birth": "10-20-2021"
    }
  ]
}
```

```Get user by id
{
  "status": true,
  "message": "success get user data by id",
  "data": {
    "id": 1,
    "name": "Nur Afan Syarifudin",
    "email": "afan@gmail.com",
    "age": "4",
    "date_of_birth": "10-20-2021"
  }
}
```

```Create new user - Payload
{
{
  "age": "string",
  "date_of_birth": "string",
  "email": "string",
  "name": "string",
  "password": "string"
}
}
```

```Delete user
{
  "status": true,
  "message": "success delete data"
}
```

```Update user
{
  "status": true,
  "message": "success update data"
}
```

## Generate docs for swagger (Run after writing the documentation in handler)

`swag init -g main.go -td "[[,]]"`

## Formattng docs in swagger

`swag fmt`

## Running via docker

### Build docker first

`docker run --rm -it backend-api`

### Running via docker

note: there is a configuration need to be adjusted db_host using `host.docker.internal`

```
docker run --rm -it -p 3000:3000  backend-api
```

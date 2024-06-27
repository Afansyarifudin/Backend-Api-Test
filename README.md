# Backend API Golang

## Applying hot-reload

`air`

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
```

## API List

http://localhost:3000/api/v1

    User Routes :
    POST("/register") register user
    POST("/login") login
    GET("/admin/users") get all profiles
    GET("/admin/users/:id") get photo by Id
    POST("/admin/users") Create profile
    PUT("/admin/users/:id") update profile
    DELETE("/admin/users/:id") delete profile

## Api docs Swagger

    GET("/docs/swagger/index.html#/") Swagger Docs

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

```Create new user
{
{
  "status": true,
  "message": "success create user",
  "data": {
    "id": 4,
    "name": "Nur Afan Syarifudin",
    "email": "afan109@gmail.com",
    "age": "4",
    "date_of_birth": "10-20-2021"
  }
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

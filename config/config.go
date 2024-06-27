package config

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type ProgrammingConfig struct {
	AppPort   int
	DBPort    uint16
	DBHost    string
	DBUser    string
	DBPass    string
	DBName    string
	Secret    string
	RefSecret string
}

// APP_URL=http://0.0.0.0
// APP_PORT=3000

// POSTGRES_ADDRESS=localhost
// POSTGRES_DB=db_test_backend
// POSTGRES_USER=postgres
// POSTGRES_PASSWORD=postgres
// POSTGRES_PORT=5432

// API_SECRET=verySecretTho

func InitConfig() *ProgrammingConfig {
	var res = new(ProgrammingConfig)
	res, _ = loadConfig()

	if res == nil {
		err := godotenv.Load(".env")
		if err != nil {
			logrus.Error("Error loading .env file", err)
			return nil
		}
		res, err = loadConfig()
		if err != nil {
			logrus.Error("Error loading configuration", err)
			return nil
		}
	}
	return res
}

func loadConfig() (*ProgrammingConfig, error) {
	var error error
	var res = new(ProgrammingConfig)
	var permit = true

	if val, found := os.LookupEnv("APP_PORT"); found {
		port, err := strconv.Atoi(val)
		if err != nil {
			logrus.Error("Config invalid app port value", err.Error())
			permit = false
		} else {
			res.AppPort = port
		}
	} else {
		permit = false
		error = errors.New("port undefined")
	}

	if val, found := os.LookupEnv("DB_HOST"); found {
		res.DBHost = val
	} else {
		permit = false
		error = errors.New("DB_HOST undefined")
	}

	if val, found := os.LookupEnv("DB_NAME"); found {
		res.DBName = val
	} else {
		permit = false
		error = errors.New("DB_NAME undefined")
	}

	if val, found := os.LookupEnv("DB_USER"); found {
		res.DBUser = val
	} else {
		permit = false
		error = errors.New("DB_USER undefined")
	}

	if val, found := os.LookupEnv("DB_PASS"); found {
		res.DBPass = val
	} else {
		permit = false
		error = errors.New("DB_PASS undefined")
	}

	if val, found := os.LookupEnv("SECRET"); found {
		res.Secret = val
	} else {
		permit = false
		error = errors.New("SECRET undefined")
	}

	if val, found := os.LookupEnv("REFSECRET"); found {
		res.Secret = val
	} else {
		permit = false
		error = errors.New("REFSECRET undefined")
	}

	if val, found := os.LookupEnv("DB_PORT"); found {
		port, err := strconv.Atoi(val)
		if err != nil {
			logrus.Error("Config invalid db port value", err.Error())
			permit = false
		} else {
			res.DBPort = uint16(port)
		}

	} else {
		permit = false
		error = errors.New("port undefined")
	}

	if !permit {
		return nil, error
	}

	return res, nil

}

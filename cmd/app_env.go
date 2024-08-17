package cmd

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"strings"
)

type AppEnvironment string

const (
	EnvDevelopment AppEnvironment = "development"
	EnvProduction  AppEnvironment = "production"
)

type AppEnv struct {
	Port               uint16
	DBDsn              string
	AppEnvironment     AppEnvironment
	ApplicationName    string
	ApplicationVersion string
	JwtSecret          string
}

func loadEnv(filenames ...string) (*AppEnv, error) {
	//TODO refacto to use functions to load env instead of else if everywhere
	if err := godotenv.Load(filenames...); err != nil {
		return nil, err
	}
	port := uint16(8080)
	dbDsn := ""
	appEnvironment := EnvDevelopment
	jwtSecret := ""

	//Load port number from environment variable
	if portStr, ok := os.LookupEnv("PORT"); ok {
		if parsedPort, err := strconv.ParseInt(portStr, 10, 16); err != nil {
			return nil, errors.New(fmt.Sprintf("'%s' is not a valid port number, setting to default %d", portStr, port))
		} else {
			port = uint16(parsedPort)
		}
	}
	if dbDsnEnv, ok := os.LookupEnv("DB_DSN"); ok {
		dbDsn = strings.TrimSpace(dbDsnEnv)
	} else {
		return nil, errors.New("DB_DSN is not set")
	}

	if jwtSecretEnv, ok := os.LookupEnv("JWT_SECRET"); ok {
		jwtSecret = jwtSecretEnv
	} else {
		return nil, errors.New("JWT_SECRET")
	}

	if appEnv, ok := os.LookupEnv("ENVIRONMENT"); ok {
		//loop over AppEnvironment constants
		isEnvValid := false
		for _, env := range []AppEnvironment{EnvDevelopment, EnvProduction} {
			if appEnv == string(env) {
				appEnvironment = env
				isEnvValid = true
				break
			}
		}
		if !isEnvValid {
			return nil, errors.New(fmt.Sprintf("Invalid environment '%s'", appEnv))
		}
	}

	return &AppEnv{
		Port:               port,
		DBDsn:              dbDsn,
		AppEnvironment:     appEnvironment,
		ApplicationName:    os.Getenv("APPLICATION_NAME"),
		ApplicationVersion: os.Getenv("APPLICATION_VERSION"),
		JwtSecret:          jwtSecret,
	}, nil
}

package cmd

import (
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"strings"
)

func setupPostgresqlDB(env AppEnv) (db *gorm.DB, err error) {
	db, err = gorm.Open(postgres.Open(env.DBDsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if host, err := getHostFromConnString(env.DBDsn); err != nil {
		return nil, errors.New("couldn't find host in connection string")
	} else {
		log.Printf("Connected to %s", host)
	}
	// get
	return db, nil
}

func getHostFromConnString(connStr string) (string, error) {
	// Split the connection string into key-value pairs
	pairs := strings.Split(connStr, " ")

	// Iterate through each pair to find the "host" key
	for _, pair := range pairs {
		keyValue := strings.SplitN(pair, "=", 2)
		if len(keyValue) == 2 && keyValue[0] == "host" {
			return keyValue[1], nil
		}
	}

	// Return an api_error if the "host" key is not found
	return "", fmt.Errorf("host not found in connection string")
}

package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"fmt"
	"errors"
)

var (
	ErrEmptyDSN = errors.New("DSN is empty")
	ErrInvalidDSN = errors.New("invalid DSN")
)

func PostgresDB(dsn string) (*gorm.DB, error) {
	if dsn == "" {
		return nil, ErrEmptyDSN
	}

	// Validate DSN format
	if !isValidDSN(dsn) {
		return nil, ErrInvalidDSN
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Customize GORM configuration here
		// For example, enable logging
		// Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}

func isValidDSN(dsn string) bool {
	// Implement DSN validation logic here
	// For example, check if the DSN contains the required parameters
	// (e.g. host, port, user, password, dbname)
	if dsn == "" {
		return false
	}

	// Split the DSN by spaces
	dsnParts := strings.Split(dsn, " ")

	// Check if the DSN contains the required parameters
	requiredParams := []string{"host", "port", "user", "password", "dbname"}
	for _, param := range requiredParams {
		if !contains(dsnParts, param) {
			return false
		}
	}

	return true
}

func contains(arr []string, element string) bool {
	for _, e := range arr {
		if e == element {
			return true
		}
	}
	return false
}

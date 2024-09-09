package models

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func SetupDBFromEnv(schemaName string) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s search_path=%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOSTNAME"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_SSLMODE"),
		schemaName,
	)

	return SetupDB(dsn, schemaName)
}

func SetupDB(dsn, schemaName string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: schemaName + ".",
		},
	})
	if err != nil {
		return nil, err
	}

	if err := db.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", schemaName)).Error; err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(All()...); err != nil {
		return nil, err
	}

	return db, nil
}

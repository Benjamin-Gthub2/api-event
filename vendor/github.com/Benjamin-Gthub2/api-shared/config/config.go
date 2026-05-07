/*
 * File: config.go
 * Author: joel
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains struct for the configuration.
 *
 * Last Modified: 2023-09-08
 */

package config

type Configuration struct {
	ServerPort        string   `envconfig:"SERVER_PORT"`
	StoragePath       string   `envconfig:"STORAGE_PATH"`
	StoragePathReport string   `envconfig:"STORAGE_PATH_REPORT"`
	WkhtmltopdfPath   string   `envconfig:"WKHTMLTOPDF_PATH"`
	Project           string   `envconfig:"PROJECT"`
	Env               string   `envconfig:"ENV"`
	LoggingUrl        string   `envconfig:"LOGGING_URL"`
	DB                DB       `envconfig:"DB"`
	JWT               JWT      `envconfig:"JWT"`
	KeyUsers          KeyUsers `envconfig:"KEY_USERS"`
}

type DB struct {
	DbDriver   string `envconfig:"DB_DRIVER"`
	DbDatabase string `envconfig:"DB_DATABASE"`
	DbHost     string `envconfig:"DB_HOST"`
	DbPassword string `envconfig:"DB_PASSWORD"`
	DbPort     string `envconfig:"DB_PORT"`
	DbUsername string `envconfig:"DB_USERNAME"`
}

type JWT struct {
	JwtSecret string `envconfig:"JWT_SECRET"`
}

type KeyUsers struct {
	KeyUsers string `envconfig:"KEY_USERS"`
}

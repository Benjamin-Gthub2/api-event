/*
 * File: connect.go
 * Author: patrick
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This is the entry point for the application.
 *
 * Last Modified: 2023-08-02
 */

package db

import "github.com/redis/go-redis/v9"

var (
	Database *DB
	Cache    *redis.Client
)

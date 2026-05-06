/*
 * File: tools.go
 * Author: joel
 * Copyright: 2023, Smart Cities Peru.
 * License: MIT
 *
 * Purpose:
 * This file contains the functions to connect to the database.
 *
 * Last Modified: 2023-09-08
 */

package db

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"

	"github.com/Benjamin-Gthub2/api-shared/config"
	errorDomain "github.com/Benjamin-Gthub2/api-shared/error-core/domain"
)

//go:embed sql/mysql/get_schema_by_tenant_id.sql
var QueryGetSchemaByTenantId string

//go:embed sql/postgres/get_schema_by_tenant_id.sql
var QueryGetSchemaByTenantIdPostgres string

const AmericaLima = "America%2FLima"
const MysqlUriFormat = "%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=%s&multiStatements=true"
const MysqlUriFormatWithoutTimezone = "%s:%s@tcp(%s:%s)/%s?parseTime=true&multiStatements=true"

// const PostgresUriFormat = "host=%s port=%s user=%s dbname=%s sslmode=disable password=%s"
const PostgresUriFormat = "host=%s port=%s user=%s dbname=%s sslmode=disable password=%s options='-c search_path=%s'"

type DB struct {
	TenantDB       *sql.DB
	ClientSchemaDB map[string]*sql.DB
}

func StartMysql(cfg config.Configuration, timeZone *string) (err error) {
	tenantUri := ""
	if timeZone != nil {
		tenantUri = fmt.Sprintf(
			MysqlUriFormat,
			cfg.DB.DbUsername,
			cfg.DB.DbPassword,
			cfg.DB.DbHost,
			cfg.DB.DbPort,
			"db_tenant",
			*timeZone,
		)
	} else {
		tenantUri = fmt.Sprintf(
			MysqlUriFormatWithoutTimezone,
			cfg.DB.DbUsername,
			cfg.DB.DbPassword,
			cfg.DB.DbHost,
			cfg.DB.DbPort,
			"db_tenant",
		)
	}
	if Database == nil {
		Database = &DB{}
	}
	Database.TenantDB, err = ConnectMySQL(tenantUri)
	if err != nil {
		return err
	}
	Database.ClientSchemaDB = make(map[string]*sql.DB)
	return nil
}

func StartPostgres(cfg config.Configuration) (err error) {
	tenantUri := fmt.Sprintf(
		PostgresUriFormat,
		cfg.DB.DbHost,
		cfg.DB.DbPort,
		cfg.DB.DbUsername,
		"db_tenant",
		cfg.DB.DbPassword,
		"db_tenant",
	)
	if Database == nil {
		Database = &DB{}
	}
	Database.TenantDB, err = ConnectPostgres(tenantUri)
	if err != nil {
		return err
	}
	Database.ClientSchemaDB = make(map[string]*sql.DB)
	return nil
}

func AddClientSchemaDB(xTenantId string, client *sql.DB) {
	if Database == nil {
		Database = &DB{}
		Database.ClientSchemaDB = make(map[string]*sql.DB)
	}
	Database.ClientSchemaDB[xTenantId] = client
}

func ClientDB(ctx context.Context) (*sql.DB, *string, error) {
	xTenantIdTmp := ctx.Value("xTenantId")
	xTenantId, hasXTenantId := xTenantIdTmp.(string)
	err := errorDomain.NewUnauthorizedErr().SetLayer(errorDomain.Infra)
	if !hasXTenantId {
		return nil, nil, err.Clone().SetFunction("ClientDB").SetRaw(fmt.Errorf("xTenantId not found"))
	}

	_, errUuid := uuid.Parse(xTenantId)
	if errUuid == nil {
		if _, hasSchema := Database.ClientSchemaDB[xTenantId]; hasSchema {
			return Database.ClientSchemaDB[xTenantId], &xTenantId, nil
		}
	}

	var schema string
	dbDriver := os.Getenv("DB_DRIVER")
	if dbDriver == "postgres" {
		errQuery := Database.TenantDB.QueryRowContext(
			ctx,
			QueryGetSchemaByTenantIdPostgres,
			xTenantId,
			xTenantId,
		).Scan(&schema, &xTenantId)
		if errQuery != nil {
			fmt.Println("err S1", errQuery)
			return nil, nil, err.Clone().SetFunction("ClientDB").SetRaw(errQuery)
		}
	} else {
		errQuery := Database.TenantDB.QueryRowContext(
			ctx,
			QueryGetSchemaByTenantId,
			xTenantId,
			xTenantId,
		).Scan(&schema, &xTenantId)
		if errQuery != nil {
			fmt.Println("err S1", errQuery)
			return nil, nil, err.Clone().SetFunction("ClientDB").SetRaw(errQuery)
		}
	}
	if schema == "" {
		return nil, nil, err.Clone().SetFunction("ClientDB").SetRaw(fmt.Errorf("schema not found"))
	}

	cfg := config.DB{
		DbDatabase: schema,
		DbDriver:   dbDriver,
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
		DbUsername: os.Getenv("DB_USERNAME"),
		DbPassword: os.Getenv("DB_PASSWORD"),
	}
	dbTimeZone := os.Getenv("DB_TIMEZONE")
	if dbTimeZone == "" {
		// default is America Lima
		dbTimeZone = AmericaLima
	} else if dbTimeZone == "NULL" {
		dbTimeZone = ""
	}
	var client *sql.DB
	var errConnect error
	if dbDriver == "postgres" {
		postgresUri := fmt.Sprintf(
			PostgresUriFormat,
			cfg.DbHost,
			cfg.DbPort,
			cfg.DbUsername,
			cfg.DbDatabase,
			cfg.DbPassword,
			cfg.DbDatabase,
		)
		client, errConnect = ConnectPostgres(postgresUri)
		if errConnect != nil {
			return nil, &xTenantId, err.Clone().SetFunction("ClientDB").SetRaw(errConnect)
		}
		// review deprecated feature
		Database.ClientSchemaDB[xTenantId] = client
	} else {
		mySqlUri := ""
		if dbTimeZone != "" {
			mySqlUri = fmt.Sprintf(
				MysqlUriFormat,
				cfg.DbUsername,
				cfg.DbPassword,
				cfg.DbHost,
				cfg.DbPort,
				cfg.DbDatabase,
				AmericaLima,
			)
		} else {
			mySqlUri = fmt.Sprintf(
				MysqlUriFormatWithoutTimezone,
				cfg.DbUsername,
				cfg.DbPassword,
				cfg.DbHost,
				cfg.DbPort,
				cfg.DbDatabase,
			)
		}

		client, errConnect = ConnectMySQL(mySqlUri)
		if errConnect != nil {
			fmt.Println("err S2", errConnect)
			return nil, &xTenantId, err.Clone().SetFunction("ClientDB").SetRaw(errConnect)
		}
		Database.ClientSchemaDB[xTenantId] = client
	}

	ctx = context.WithValue(ctx, "xTenantId", xTenantId)

	return client, &xTenantId, nil
}

// InitClients deprecate
func InitClients(cfg config.Configuration) (err error) {
	timezone := AmericaLima
	if cfg.DB.DbDriver == "postgres" {
		err = StartPostgres(cfg)
	} else {
		err = StartMysql(cfg, &timezone)
	}
	if err != nil {
		return err
	}
	// Ensure Database and ClientSchemaDB are initialized
	if Database == nil {
		Database = &DB{}
		Database.ClientSchemaDB = make(map[string]*sql.DB)
	}
	if Database.ClientSchemaDB == nil {
		Database.ClientSchemaDB = make(map[string]*sql.DB)
	}
	var client *sql.DB
	if cfg.DB.DbDriver == "postgres" {
		postgresUri := fmt.Sprintf(
			PostgresUriFormat,
			cfg.DB.DbHost,
			cfg.DB.DbPort,
			cfg.DB.DbUsername,
			cfg.DB.DbDatabase,
			cfg.DB.DbPassword,
			cfg.DB.DbDatabase,
		)
		client, err = ConnectPostgres(postgresUri)
		if err != nil {
			return err
		}
		// review deprecated feature
		Database.ClientSchemaDB[cfg.DB.DbDatabase] = client
	} else {
		mySqlUri := fmt.Sprintf(
			MysqlUriFormat,
			cfg.DB.DbUsername,
			cfg.DB.DbPassword,
			cfg.DB.DbHost,
			cfg.DB.DbPort,
			cfg.DB.DbDatabase,
			AmericaLima,
		)
		client, err = ConnectMySQL(mySqlUri)
		if err != nil {
			return err
		}
		// review deprecated feature
		Database.ClientSchemaDB[cfg.DB.DbDatabase] = client
	}
	return nil
}

// InitClients deprecate
func InitClientsWithoutTimeZone(cfg config.Configuration) (err error) {
	err = StartMysql(cfg, nil)
	if err != nil {
		return err
	}
	var (
		MySqlUri = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?parseTime=true&multiStatements=true",
			cfg.DB.DbUsername,
			cfg.DB.DbPassword,
			cfg.DB.DbHost,
			cfg.DB.DbPort,
			cfg.DB.DbDatabase,
		)
	)
	var client *sql.DB
	client, err = ConnectMySQL(MySqlUri)
	if err != nil {
		return err
	}
	// review deprecated feature
	Database.ClientSchemaDB[cfg.DB.DbDatabase] = client
	return nil
}

func ConnectMySQL(uri string) (client *sql.DB, err error) {
	client, err = sql.Open("mysql", uri)
	if err != nil {
		return nil, err
	}
	client.SetConnMaxLifetime(time.Minute * 6)
	client.SetMaxOpenConns(200)
	client.SetMaxIdleConns(200)
	return client, nil
}

func Disconnect() {
	if Database != nil {
		if Database.TenantDB != nil {
			Database.TenantDB.Close()
		}
		if Database.ClientSchemaDB != nil {
			for _, database := range Database.ClientSchemaDB {
				if database != nil {
					database.Close()
				}
			}
		}
	}
}

func ConnectSqlServer(uri string) (client *sql.DB, err error) {
	client, err = sql.Open("mssql", uri)
	if err != nil {
		return nil, err
	}
	client.SetConnMaxLifetime(time.Minute * 6)
	client.SetMaxOpenConns(200)
	client.SetMaxIdleConns(200)
	return client, nil
}

func ConnectPostgres(uri string) (client *sql.DB, err error) {
	client, err = sql.Open("postgres", uri)
	if err != nil {
		return nil, err
	}
	client.SetConnMaxLifetime(time.Minute * 6)
	client.SetMaxOpenConns(200)
	client.SetMaxIdleConns(200)
	return client, nil
}

func ConnectRedis() (err error) {
	redisURI := fmt.Sprintf(
		"redis://%s:%s@%s:%s",
		"",
		os.Getenv("REDIS_PASSWORD"),
		os.Getenv("REDIS_HOST"),
		os.Getenv("REDIS_PORT"),
	)
	opt, err := redis.ParseURL(redisURI)
	if err != nil {
		panic(err)
	}
	Cache = redis.NewClient(opt)

	var ctx = context.Background()
	_, err = Cache.Ping(ctx).Result()
	if err != nil {
		return fmt.Errorf("Error connecting to Redis: %v", err)
	}
	fmt.Println("Connected to Redis")

	return nil
}

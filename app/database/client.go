package database

import (
	"fmt"
	"strconv"
	"time"

	"github.com/LOCNNIL/golang-rinha-api/app/environment"
	"github.com/LOCNNIL/golang-rinha-api/app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConnection struct {
	host          string
	port          string
	user          string
	password      string
	database_name string
	max_retries   int
	internal      int
}

func Migrate(database_connection *gorm.DB) {
	database_connection.AutoMigrate(
		&models.People{},
	)
}

func GetDatabaseConnectionFromEnvVars() DatabaseConnection {
	max_retries, err := strconv.Atoi(environment.GetEnvOrDie("POSTGRES_CONNECTION_RETRIES"))
	if err != nil {
		panic(err.Error())
	}

	interval, err := strconv.Atoi(environment.GetEnvOrDie("POSTGRES_RETRIES_INTERVAL"))
	if err != nil {
		panic(err.Error())
	}

	return DatabaseConnection{
		host:          environment.GetEnvOrDie("POSTGRES_HOST"),
		port:          environment.GetEnvOrDie("POSTGRES_PORT"),
		user:          environment.GetEnvOrDie("POSTGRES_USER"),
		password:      environment.GetEnvOrDie("POSTGRES_PASSWORD"),
		database_name: environment.GetEnvOrDie("POSTGRES_DB"),
		max_retries:   max_retries,
		internal:      interval,
	}
}

func ParseDatabaseStruct2ConnectionString(conn DatabaseConnection) string {
	conn_string := "host=" + conn.host + " " +
		"port=" + conn.port + " " +
		"user=" + conn.user + " " +
		"password=" + conn.password + " " +
		"dbname=" + conn.database_name + " " +
		"sslmode=disable"
	return conn_string
}

func triesToConnectWithDatabase(connection_string string, db_struct DatabaseConnection) (*gorm.DB, error) {
	var conn *gorm.DB
	var err error
	for i := 0; i < db_struct.max_retries; i++ {
		conn, err = gorm.Open(postgres.Open(connection_string), &gorm.Config{})
		if err == nil {
			return conn, nil
		}
		time.Sleep(time.Second * time.Duration(db_struct.internal))
	}

	return nil, err
}

func CreateConnection() *gorm.DB {
	database_conn := GetDatabaseConnectionFromEnvVars()
	dsn := ParseDatabaseStruct2ConnectionString(database_conn)
	connection, err := triesToConnectWithDatabase(dsn, database_conn)

	if err != nil {
		msg := fmt.Sprintf("[Error] opening database: %s", err.Error())
		panic(msg)
	}
	return connection
}

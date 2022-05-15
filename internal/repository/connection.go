package repository

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
)

// Db holds the databse connection pool
type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

func GetConnection() *sql.DB {
	return dbConn.SQL
}

// Connect created database pool for database
func Connect(dsn string) (*sql.DB, error) {
	log.Println("Connecting to database")
	conn, err := NewDatabase(dsn)
	if err != nil {
		panic(err)
	}

	dbConn.SQL = conn

	err = TestConnection(conn)
	if err != nil {
		return nil, err
	}
	log.Println("Connection success")
	return dbConn.SQL, nil
}

// TestConnection tries to ping database
func TestConnection(conn *sql.DB) error {
	log.Println("Testing connection")
	err := conn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

// NewDatabase creates new database
func NewDatabase(dsn string) (*sql.DB, error) {
	conn, err := sql.Open("postgres", dsn)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return conn, nil
}

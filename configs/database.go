package configs

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func InitDB() *sql.DB {
	conn := connectToDB()
	if conn == nil {
		log.Panic("error connecting to the database")
	}

	return conn
}

func connectToDB() *sql.DB {
	counts := 0

	dsn := os.Getenv("DSN")

	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Printf("postgres didn't responded. Error %s", err)
		} else {
			log.Println("connected to database")
			return connection
		}

		if counts > 10 {
			return nil
		}

		log.Print("backing off for 1 second")
		time.Sleep(time.Second * 1)
		continue
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

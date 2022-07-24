package main

import (
	"authentication/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)


const webPort = "80"
var counts int64

type Config struct {
	Db *sql.DB
	Models data.Models
}

func main() {
	log.Println("Starting authentication service")

	// connect to db
	conn := connectToDb()
	if conn == nil {
		log.Panic("Failed to connect to DB!")
	}

	// set up config
	app := Config{
		Db: conn,
		Models: data.New(conn),
	}

	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()

	if err != nil {
		log.Panic(err)
	}
}

func openDb(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("pgx",connectionString)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectToDb() *sql.DB {
	connString := os.Getenv("DB_CONNECTIONSTRING")

	for {
		conn, err := openDb(connString)

		if err != nil {
			log.Println("Postgres not yet ready")
			counts++
		}else{
			log.Println("Connected to Postgres")
			return conn
		}

		if counts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Will attempt connection after 2 seconds")
		time.Sleep(2 * time.Second)
		continue
	}
}
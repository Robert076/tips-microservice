package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Robert076/tips-microservice/tip"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func GetTipsFromDatabase() []tip.Tip {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	// nolint:errcheck
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT Id, Text FROM Tips")

	if err != nil {
		panic(err)
	}

	// nolint:errcheck
	defer rows.Close()

	var tips []tip.Tip

	for rows.Next() {
		var t tip.Tip
		err := rows.Scan(&t.Id, &t.Text)
		if err != nil {
			panic(err)
		}
		tips = append(tips, t)
	}

	if err = rows.Err(); err != nil {
		panic(err)
	}

	return tips
}

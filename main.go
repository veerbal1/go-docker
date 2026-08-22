package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	db, err := sql.Open("pgx", "postgres://postgres:12345678@postgres-db:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal("err:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping db: ", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Got request, bye")
		fmt.Fprintln(w, "hello")
	})

	fmt.Println("Listening on port: 8080")
	http.ListenAndServe(":8080", nil)
}

// docker build -t go-docker-app .
// docker stop relaxed_agnesi
// docker rm relaxed_agnesi
// docker run --network mynet --name go-app -p 8080:8080 -d go-docker-app
// docker logs go-app

package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load("/home/vagrant/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"), os.Getenv("MOVIES_DB"))
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	handler := &InventoryHandler{db: db}
	r := mux.NewRouter()
	r.HandleFunc("/api/movies", handler.getMovies).Methods("GET")
	r.HandleFunc("/api/movies", handler.createMovie).Methods("POST")
	r.HandleFunc("/api/movies", handler.deleteAllMovies).Methods("DELETE")
	r.HandleFunc("/api/movies/{id}", handler.getMovie).Methods("GET")
	r.HandleFunc("/api/movies/{id}", handler.updateMovie).Methods("PUT")
	r.HandleFunc("/api/movies/{id}", handler.deleteMovie).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":"+os.Getenv("INVENTORY_PORT"), r))
}
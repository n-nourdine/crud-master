package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"net/http"
	"strconv"
)

type InventoryHandler struct {
	db *sqlx.DB
}

func (h *InventoryHandler) getMovies(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("title")
	var movies []Movie
	var err error
	if query != "" {
		err = h.db.Select(&movies, "SELECT * FROM movies WHERE title ILIKE $1", "%"+query+"%")
	} else {
		err = h.db.Select(&movies, "SELECT * FROM movies")
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(movies)
}

func (h *InventoryHandler) getMovie(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var movie Movie
	err := h.db.Get(&movie, "SELECT * FROM movies WHERE id=$1", id)
	if err != nil {
		http.Error(w, "Movie not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(movie)
}

func (h *InventoryHandler) createMovie(w http.ResponseWriter, r *http.Request) {
	var movie Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	err := h.db.QueryRow("INSERT INTO movies (title, description) VALUES ($1, $2) RETURNING id", movie.Title, movie.Description).Scan(&movie.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(movie)
}

func (h *InventoryHandler) updateMovie(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var movie Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	_, err := h.db.Exec("UPDATE movies SET title=$1, description=$2 WHERE id=$3", movie.Title, movie.Description, id)
	if err != nil {
		http.Error(w, "Movie not found", http.StatusNotFound)
		return
	}
	movie.ID = id
	json.NewEncoder(w).Encode(movie)
}

func (h *InventoryHandler) deleteMovie(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	result, err := h.db.Exec("DELETE FROM movies WHERE id=$1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		http.Error(w, "Movie not found", http.StatusNotFound)
		return
	}
	w.Write([]byte("Movie deleted"))
}

func (h *InventoryHandler) deleteAllMovies(w http.ResponseWriter, r *http.Request) {
	_, err := h.db.Exec("DELETE FROM movies")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("All movies deleted"))
}
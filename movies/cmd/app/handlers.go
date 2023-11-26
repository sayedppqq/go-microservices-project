package app

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sayedppqq/go-microservices-project/movies/pkg/models"
	"net/http"
)

func (app *application) all(w http.ResponseWriter, r *http.Request) {
	// Get all movie stored
	movies, err := app.movies.GetAllMovies()
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Convert movie list into json encoding
	b, err := json.Marshal(movies)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.infoLog.Println("Movies have been listed")

	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) findByID(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	// Find movie by id
	m, err := app.movies.GetMovieByID(id)
	if err != nil {
		if err.Error() == "ErrNoDocuments" {
			app.infoLog.Println("Movie not found")
			return
		}
		// Any other error will send an internal server error
		app.serverError(w, err)
		return
	}

	// Convert movie to json encoding
	b, err := json.Marshal(m)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.infoLog.Println("Have been found a movie")

	// Send response back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) insert(w http.ResponseWriter, r *http.Request) {
	// Define movie model
	var m models.Movie
	// Get request information
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Insert new movie
	insertResult, err := app.movies.InsertNewMovie(m)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.infoLog.Printf("New movie have been created, id=%s", insertResult.InsertedID)
}

func (app *application) delete(w http.ResponseWriter, r *http.Request) {
	// Get id from incoming url
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println("dlt calll")

	// Delete movie by id
	deleteResult, err := app.movies.DeleteMovieByID(id)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.infoLog.Printf("Have been eliminated %d movie(s)", deleteResult.DeletedCount)
}

package main

import (
	"fmt"
	"net/http"
	"time"

	"greenlight.mustik.net/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}

func (app *application) showMovieHandler(w http.ResponseWriter, req *http.Request) {
	id, err := app.readIDParam(req)
	if err != nil {
		http.NotFound(w, req)
		return
	}

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Harry Potter",
		Runtime:   144,
		Genres:    []string{"drama", "romance", "fantasy"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.logger.Print(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}

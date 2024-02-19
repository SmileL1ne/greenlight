package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"greenlight.mustik.net/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, req *http.Request) {
	var input struct {
		Title   string   `json:"title"`
		Year    int32    `json:"year"`
		Runtime int32    `json:"runtime"`
		Genres  []string `json:"genres"`
	}

	if err := json.NewDecoder(req.Body).Decode(&input); err != nil {
		app.errorResponse(w, req, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showMovieHandler(w http.ResponseWriter, req *http.Request) {
	id, err := app.readIDParam(req)
	if err != nil {
		app.notFoundResponse(w, req)
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
		app.serverErrorResponse(w, req, err)
	}
}

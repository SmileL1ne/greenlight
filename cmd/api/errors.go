package main

import (
	"fmt"
	"net/http"
)

func (app *application) logError(req *http.Request, err error) {
	app.logger.PrintError(err, map[string]string{
		"request_method": req.Method,
		"request_url":    req.URL.String(),
	})
}

func (app *application) errorResponse(w http.ResponseWriter, req *http.Request, status int, message any) {
	env := envelope{"error": message}

	if err := app.writeJSON(w, status, env, nil); err != nil {
		app.logError(req, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *application) serverErrorResponse(w http.ResponseWriter, req *http.Request, err error) {
	app.logError(req, err)

	message := "the server encountered a problem and could not process your request"
	app.errorResponse(w, req, http.StatusInternalServerError, message)
}

func (app *application) notFoundResponse(w http.ResponseWriter, req *http.Request) {
	message := "the requested resource could not be found"
	app.errorResponse(w, req, http.StatusNotFound, message)
}

func (app *application) methodNotAllowedResponse(w http.ResponseWriter, req *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", req.Method)
	app.errorResponse(w, req, http.StatusMethodNotAllowed, message)
}

func (app *application) badRequestResponse(w http.ResponseWriter, req *http.Request, err error) {
	app.errorResponse(w, req, http.StatusBadRequest, err.Error())
}

func (app *application) failedValidationResponse(w http.ResponseWriter, req *http.Request, errors map[string]string) {
	app.errorResponse(w, req, http.StatusUnprocessableEntity, errors)
}

func (app *application) editConflictResponse(w http.ResponseWriter, req *http.Request) {
	message := "unable to update the record due to an edit conflict, please try again"
	app.errorResponse(w, req, http.StatusConflict, message)
}

func (app *application) rateLimitExceededResponse(w http.ResponseWriter, req *http.Request) {
	message := "rate limit exceeded"
	app.errorResponse(w, req, http.StatusTooManyRequests, message)
}

package main

import (
	"net/http"
)

type envelope map[string]interface{}

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	env := envelope{"error": message}

	err := WriteJSON(w, status, env, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(500)
	}
}

func (app *application) logError(r *http.Request, err error) {
	app.logger.Println(err)
}

package main

import (
	"github.com/Hopertz/go-url-shortener/store"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
	"time"
)

type application struct {
	version int
	logger  *log.Logger
}

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	app := &application{version: 1, logger: logger}
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/", app.welcomeHandler)
	router.HandlerFunc(http.MethodPost, "/create-short-url", app.createShortUrl)
	router.HandlerFunc(http.MethodGet, "/:shorturl", app.handleShortUrlRedirect)

	store.InitializeStore()
	srv := &http.Server{
		Addr:         ":9808",
		Handler:      router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	logger.Printf("Starting server")
	err := srv.ListenAndServe()
	logger.Fatal(err)
}

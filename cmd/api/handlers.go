package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Hopertz/go-url-shortener/shortener"
	"github.com/Hopertz/go-url-shortener/store"
	"github.com/julienschmidt/httprouter"
)

type UrlCReationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}

func (app *application) welcomeHandler(w http.ResponseWriter, r *http.Request) {
	js := `{"status" :"available"}`
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(js))
}

func (app *application) createShortUrl(w http.ResponseWriter, r *http.Request) {
	var creationRequest UrlCReationRequest
	err := json.NewDecoder(r.Body).Decode(&creationRequest)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}
	shortUrl := shortener.GenerateShortLink(creationRequest.LongUrl, creationRequest.UserId)
	store.SaveUrlMapping(shortUrl, creationRequest.LongUrl, creationRequest.UserId)

	host := "http://localhost:9808/"
	newUrl := host + shortUrl

	js := `{"message": "short url created successfully", "url": %q}`
	js = fmt.Sprintf(js, newUrl)
	fmt.Fprint(w, "%+V\n", js)
}

func (app *application) handleShortUrlRedirect(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	shortUrl := params.ByName("shorturl")
	initialUrl := store.RetrieveInitialUrl(shortUrl)
	http.Redirect(w, r, initialUrl, http.StatusFound)

}

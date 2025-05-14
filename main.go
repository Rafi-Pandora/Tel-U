package main

import (
	"fmt"
	"net/http"
	"server/controller"
	"server/models"
)

func main() {
	const port = "8080"
	var listFilmArr []*models.Film

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/movie", http.StatusPermanentRedirect)
	})

	http.HandleFunc("/movie", func(w http.ResponseWriter, r *http.Request) {
		controller.MainPageHandler(w, r, &listFilmArr)
	})

	http.HandleFunc("/movie/add", func(w http.ResponseWriter, r *http.Request) {
		controller.AddFilm(w, r, &listFilmArr)
	})

	http.HandleFunc("/movi/list", func(w http.ResponseWriter, r *http.Request) {
		controller.ListFilm(w, r, &listFilmArr)
	})

	fmt.Println("Server jalan di http://localhost:" + port + "/movie")
	http.ListenAndServe(":"+port, nil)
}

package main

import (
	"fmt"
	"log"
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

	http.HandleFunc("/movie/list/sort", func(w http.ResponseWriter, r *http.Request) {
		controller.ListFilmSort(w, r, listFilmArr)
	})

	http.HandleFunc("/movie/search/result", func(w http.ResponseWriter, r *http.Request) {
		controller.SearchFilmHandler(w, r, listFilmArr)
	})

	http.HandleFunc("/movie/advncsearch", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./view/search.html")
	})

	http.HandleFunc("/movie/delete", func(w http.ResponseWriter, r *http.Request) {
		controller.DeleteFilmHandler(w, r, &listFilmArr)
	})

	http.HandleFunc("/movie/edit", func(w http.ResponseWriter, r *http.Request) {
		controller.EditFilmHandler(w, r, &listFilmArr)
	})

	http.HandleFunc("/movie/seed", func(w http.ResponseWriter, r *http.Request) {
		// if r.Method != http.MethodPost {
		// 	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		// 	return
		// }

		seedData := models.SeedFilms()
		listFilmArr = listFilmArr[:0]
		listFilmArr = append(listFilmArr, seedData...)

		log.Printf("data diinjeksi, total: %d film\n", len(listFilmArr))

		http.Redirect(w, r, "/movie", http.StatusSeeOther)
	})

	fmt.Println("Server jalan di http://localhost:" + port + "/movie")
	http.ListenAndServe(":"+port, nil)
}

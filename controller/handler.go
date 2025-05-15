package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"server/models"
	"server/utils"
	"slices"
	"strconv"
)

type ViewData struct {
	Name  string
	Films []*models.Film
}

func MainPageHandler(w http.ResponseWriter, r *http.Request, data *[]*models.Film) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Pengunjung"
	}

	funcMap := template.FuncMap{
		"add": func(a, b int) int { return a + b },
	}

	viewData := ViewData{
		Name:  name,
		Films: *data,
	}

	t, err := template.New("film.html").Funcs(funcMap).ParseFiles("./view/film.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("Template parsing error: %v", err), http.StatusInternalServerError)
		return
	}

	if err := t.Execute(w, viewData); err != nil {
		http.Error(w, fmt.Sprintf("Template execution error: %v", err), http.StatusInternalServerError)
	}
}

func AddFilm(w http.ResponseWriter, r *http.Request, data *[]*models.Film) {
	if r.Method == http.MethodGet {
		tmpl := template.Must(template.ParseFiles("./view/addFilm.html"))
		tmpl.Execute(w, nil)
	}

	if r.Method == http.MethodPost {
		judul := r.FormValue("judul")
		genre := r.FormValue("genre")
		rating := r.FormValue("rating")
		status := r.FormValue("status")

		var statusBool bool
		if status == "ditonton" || status == "Ditonton" {
			statusBool = true
		} else if status == "belum" || status == "Belum" {
			statusBool = false
		}

		ratingFloat32, err := strconv.ParseFloat(rating, 32)
		if err != nil {
			http.Error(w, "tidak dapat membaca atau mengkonversi float", http.StatusInternalServerError)
		}

		film := models.Film{
			Judul:  judul,
			Genre:  genre,
			Rating: float32(ratingFloat32),
			Status: statusBool,
		}

		*data = append(*data, &film)
		log.Println("data ditambahkan", film)

		http.Redirect(w, r, "/movie", http.StatusPermanentRedirect)
	}
}

func ListFilmSort(w http.ResponseWriter, r *http.Request, data []*models.Film) {
	sort := r.URL.Query().Get("sort")
	name := r.URL.Query().Get("name")

	//
	films := make([]*models.Film, 0, len(data))
	for _, f := range data {
		if f != nil {
			films = append(films, f)
		}
	}

	switch sort {
	case "rating_asc":
		utils.InsertionSortByRating(films, true)
	case "rating_desc":
		utils.InsertionSortByRating(films, false)
	case "title_asc":
		utils.InsertionSortByJudul(films, true)
	case "title_desc":
		utils.InsertionSortByJudul(films, false)
	default:
		http.Error(w, "Request Param Failed", http.StatusBadRequest)
	}

	funcMap := template.FuncMap{
		"add": func(a, b int) int { return a + b },
	}

	viewData := ViewData{
		Name:  name,
		Films: films,
	}

	tmpl, err := template.New("film.html").Funcs(funcMap).ParseFiles("./view/film.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("Template parsing error: %v", err), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, viewData); err != nil {
		http.Error(w, fmt.Sprintf("Template execution error: %v", err), http.StatusInternalServerError)
	}
}

func SearchFilmHandler(w http.ResponseWriter, r *http.Request, data []*models.Film) {
	genreParam := r.URL.Query().Get("genre")
	judulParam := r.URL.Query().Get("judul")

	// Copy []*Film ke []Film (tanpa pointer)
	films := make([]*models.Film, 0, len(data))
	for _, f := range data {
		if f != nil {
			films = append(films, f)
		}
	}

	var found *models.Film

	// pencarian berdasarkan genre
	if genreParam != "" {
		utils.InsertionSortByGenre(films, true)
		if _, result := utils.BinarySearchByGenre(films, genreParam); result != nil {
			found = result
		}
	}

	// pencarian berdasarkan judul
	if judulParam != "" {
		utils.InsertionSortByJudul(films, true)
		if _, result := utils.BinarySearchByJudul(films, judulParam); result != nil {
			found = result
		}
	}

	viewData := ViewData{
		Name:  "Hasil Pencarian",
		Films: []*models.Film{},
	}
	if found != nil {
		viewData.Films = append(viewData.Films, found)
	}

	funcMap := template.FuncMap{
		"add": func(a, b int) int { return a + b },
	}

	tmpl, err := template.New("film.html").Funcs(funcMap).ParseFiles("./view/film.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("Template parsing error: %v", err), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, viewData); err != nil {
		http.Error(w, fmt.Sprintf("Template execution error: %v", err), http.StatusInternalServerError)
	}
	log.Println(found)
}

func DeleteFilmHandler(w http.ResponseWriter, r *http.Request, data *[]*models.Film) {
	judul := r.URL.Query().Get("judul")
	if judul == "" {
		http.Error(w, "Param Parsing Failed", http.StatusBadRequest)
		return
	}

	utils.InsertionSortByJudul(*data, true)
	_, result := utils.BinarySearchByJudul(*data, judul)
	if result == nil {
		http.Error(w, "Film tidak ditemukan", http.StatusNotFound)
		return
	}

	idx := -1
	for i, f := range *data {
		if f == result {
			idx = i
			break
		}
	}

	if idx != -1 {
		*data = slices.Delete((*data), idx, idx+1)
		log.Println(data, "berhasil dihapus")
	}

	http.Redirect(w, r, "/movie", http.StatusSeeOther)
}

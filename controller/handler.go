package controller

import (
	"fmt"
	"log"
	"net/http"
	"server/models"
	"server/utils"
	"slices"
)

func MainPageHandler(w http.ResponseWriter, r *http.Request, data *[]*models.Film) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Pengunjung"
	}

	viewData := utils.ViewData{
		Name:  name,
		Films: *data,
	}

	t, err := utils.ParseTemplate("film.html", "./view/film.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("Template parsing error: %v", err), http.StatusInternalServerError)
		return
	}

	if err := t.Execute(w, viewData); err != nil {
		http.Error(w, fmt.Sprintf("Template execution error: %v", err), http.StatusInternalServerError)
	}
}

func AddFilm(w http.ResponseWriter, r *http.Request, data *[]*models.Film) {
	viewData := utils.FilmFormData{
		FormTitle:   "Tambah Film",
		FormAction:  "/movie/add",
		SubmitLabel: "Tambah",
		Film:        nil,
	}

	switch r.Method {
	case http.MethodGet:
		tmpl, err := utils.ParseTemplate("filmForm.html", "./view/filmForm.html")
		if err != nil {
			http.Error(w, "Template parsing error", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, viewData)

	case http.MethodPost:
		judul := r.FormValue("judul")
		genre := r.FormValue("genre")
		statusStr := r.FormValue("status")

		statusBool := statusStr == "ditonton" || statusStr == "Ditonton"

		ratingInput := r.FormValue("rating_input")
		ratingSelect := r.FormValue("rating_select")

		finalRating, err := utils.RatingFinal(ratingInput, ratingSelect)

		if err != nil {
			http.Error(w, "Tidak dapat membaca atau mengkonversi rating", http.StatusBadRequest)
			return
		}

		film := &models.Film{
			Judul:  judul,
			Genre:  genre,
			Rating: float32(finalRating),
			Status: statusBool,
		}

		*data = append(*data, film)
		log.Println("Data ditambahkan:", film)
		http.Redirect(w, r, "/movie", http.StatusSeeOther)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func ListFilmSort(w http.ResponseWriter, r *http.Request, data []*models.Film) {
	sortParam := r.URL.Query().Get("sort")
	name := r.URL.Query().Get("name")

	films := utils.FilterNonNilFilms(data)

	switch sortParam {
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
		return
	}

	viewData := utils.ViewData{Name: name, Films: films}

	tmpl, err := utils.ParseTemplate("film.html", "./view/film.html")
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

	films := utils.FilterNonNilFilms(data)

	var found *models.Film

	if genreParam != "" {
		utils.InsertionSortByGenre(films, true)
		if _, res := utils.BinarySearchByGenre(films, genreParam); res != nil {
			found = res
		}
	}

	if judulParam != "" {
		utils.InsertionSortByJudul(films, true)
		if _, res := utils.BinarySearchByJudul(films, judulParam); res != nil {
			found = res
		}
	}

	viewData := utils.ViewData{Name: "Hasil Pencarian", Films: []*models.Film{}}
	if found != nil {
		viewData.Films = append(viewData.Films, found)
	}

	tmpl, err := utils.ParseTemplate("film.html", "./view/film.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("Template parsing error: %v", err), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, viewData); err != nil {
		http.Error(w, fmt.Sprintf("Template execution error: %v", err), http.StatusInternalServerError)
	}
	log.Println("Search result:", found)
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

	idx := utils.FindFilmIndex(*data, result)
	if idx != -1 {
		*data = slices.Delete(*data, idx, idx+1)
		log.Println("Data berhasil dihapus:", *data)
	}

	http.Redirect(w, r, "/movie", http.StatusSeeOther)
}

func EditFilmHandler(w http.ResponseWriter, r *http.Request, data *[]*models.Film) {
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

	idx := utils.FindFilmIndex(*data, result)
	if idx == -1 {
		http.Error(w, "Film tidak ditemukan di slice", http.StatusInternalServerError)
		return
	}

	switch r.Method {
	case http.MethodGet:
		viewData := utils.FilmFormData{
			FormTitle:   "Edit Film",
			FormAction:  "/movie/edit?judul=" + result.Judul,
			SubmitLabel: "Simpan",
			Film:        result,
		}

		tmpl, err := utils.ParseTemplate("filmForm.html", "./view/filmForm.html")
		if err != nil {
			http.Error(w, "Template error", http.StatusInternalServerError)
			return
		}

		if err := tmpl.Execute(w, viewData); err != nil {
			http.Error(w, "Template execute error", http.StatusInternalServerError)
		}

	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Form parsing error", http.StatusBadRequest)
			return
		}

		newJudul := r.FormValue("judul")
		newGenre := r.FormValue("genre")
		newStatus := r.FormValue("status") == "ditonton"

		ratingInput := r.FormValue("rating_input")
		ratingSelect := r.FormValue("rating_select")

		finalRating, err := utils.RatingFinal(ratingInput, ratingSelect)
		if err != nil {
			http.Error(w, "Invalid rating", http.StatusBadRequest)
			return
		}

		(*data)[idx].Judul = newJudul
		(*data)[idx].Genre = newGenre
		(*data)[idx].Rating = float32(finalRating)
		(*data)[idx].Status = newStatus

		log.Println("Data berhasil diubah:", *data)
		http.Redirect(w, r, "/movie/list", http.StatusSeeOther)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

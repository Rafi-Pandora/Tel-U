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
		Name:      name,
		FormLabel: "Belum ada data film.",
		Films:     *data,
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
	ratingParam := r.URL.Query().Get("rating")

	films := utils.FilterNonNilFilms(data)

	var found []*models.Film
	var formLabel, textResult string
	var err error

	if genreParam != "" {
		if found = utils.SequentialSearchByGenre(films, genreParam); found != nil {
			textResult = "Hasil Pencarian genre " + genreParam
		} else {
			formLabel = "film dengan genre " + genreParam + " tidak ditemukan"
		}
	}

	if ratingParam != "" {
		found, err = utils.SequentialSearchByRating(films, ratingParam)

		if err != nil {
			http.Error(w, "parsing error", http.StatusInternalServerError)
		}

		if found != nil {
			textResult = "Hasil Pencarian film dengan rating " + ratingParam
		} else {
			formLabel = "film dengan film dengan rating " + ratingParam + " tidak ditemukan"
		}
	}

	if judulParam != "" {
		utils.InsertionSortByJudul(films, true)
		if indexFilm, res := utils.BinarySearchByJudul(films, judulParam); res != nil {
			found = append(found, res)
			log.Println("Search result: ", res, ", urutan index: ", indexFilm)
			textResult = "Hasil Pencarian judul " + judulParam
		} else {
			formLabel = "film dengan judul " + judulParam + " tidak ditemukan"
		}
	}

	viewData := utils.ViewData{
		Name:      textResult,
		FormLabel: formLabel,
		Films:     []*models.Film{},
	}
	if found != nil {
		viewData.Films = found
		for _, f := range found {
			log.Printf("Search result: %+v\n", *f)
		}
	}

	tmpl, err := utils.ParseTemplate("film.html", "./view/film.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("Template parsing error: %v", err), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, viewData); err != nil {
		http.Error(w, fmt.Sprintf("Template execution error: %v", err), http.StatusInternalServerError)
	}
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

		log.Printf("data dengan \njudul: %s\nGenre: %s\nRating: %.1f\nStatus: %t \nBerhasil diubah.\n", newJudul, newGenre, float32(finalRating), newStatus)

		http.Redirect(w, r, "/movie/list", http.StatusSeeOther)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

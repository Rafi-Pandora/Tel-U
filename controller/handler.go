package controller

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"server/models"
	"strconv"
	"time"
)

func MainPageHandler(w http.ResponseWriter, r *http.Request, data *[]*models.Film) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Pengunjung"
	}

	funcMap := template.FuncMap{
		"add": func(a, b int) int { return a + b },
	}

	type ViewData struct {
		Name  string
		Films []*models.Film
	}

	viewData := ViewData{
		Name:  name,
		Films: *data,
	}

	t, err := template.New("film.html").Funcs(funcMap).ParseFiles("./view/film.html")
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}

	if err := t.Execute(w, viewData); err != nil {
		http.Error(w, "Template execute error", http.StatusInternalServerError)
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

		time.Sleep(3 * time.Second)
		http.Redirect(w, r, "/movie", http.StatusPermanentRedirect)
	}
}

func ListFilm(w http.ResponseWriter, r *http.Request, data *[]*models.Film) {
	// Cek apakah client minta JSON (misal: lewat query ?format=json)
	format := r.URL.Query().Get("format")

	if format == "json" {
		if r.Method != http.MethodPost {
			http.Error(w, "Harus Menggunakan POST", http.StatusMethodNotAllowed)
			return
		}
		// Kirim response JSON saja
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
		return
	}

	funcMap := template.FuncMap{
		"add": func(a, b int) int { return a + b },
	}

	t, err := template.New("userlist").Funcs(funcMap).ParseFiles("./view/film.html")
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}

	t.Execute(w, data)
}

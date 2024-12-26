package main

import (
	"log"
	"net/http"
)

func main() {
	//port server
	const port string = ":8080"

	//inisialisasi database untuk rounting
	db, err := InsialisasiDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//variabel unutk menyimpan array hasil dari fungsi
	nama, nim := LihatSemuaUserDiDatabase(db)

	//===============================================================================//

	//endpoint handler
	http.HandleFunc("/get", GetCookie)

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		LoginHandler(w, r, db)
	})

	http.HandleFunc("/del", func(w http.ResponseWriter, r *http.Request) {
		HapusMahasiswa(w, r, db)
	})

	http.HandleFunc("/tambah", func(w http.ResponseWriter, r *http.Request) {
		err := TambahMahasiswa(w, r, db)
		if err != nil {
			log.Fatal("data gagal ditambahkan")
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		HandlingDashboard(w, r, &nama, &nim)
	})

	println("server berjalan pada http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}

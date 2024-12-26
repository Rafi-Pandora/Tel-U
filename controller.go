package main

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

var percobaan uint8 = 0
var listValueCookie = &[3]string{"true", "false", ""}
var listNamaCookie = &[2]string{"login", "blokir"}

func buatCookie(w http.ResponseWriter, _ *http.Request, NamaCookie string, NilaiCookie string) {
	cookie := &http.Cookie{
		Name:     NamaCookie,
		Value:    NilaiCookie,
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add(60 * time.Second),
	}
	http.SetCookie(w, cookie)
}

func cekCookie(_ http.ResponseWriter, r *http.Request, NamaCookie string, ValueArr *[]string) string {
	var adaCookie bool
	cekCookie, err := r.Cookie(NamaCookie)
	if err != nil || !adaCookie {
		return "cookie tidak ditemukan"
	}
	for i := 0; i < len(*ValueArr); i++ {
		if cekCookie.Value == (*ValueArr)[i] {
			adaCookie = true
		}
	}
	return cekCookie.Value
}

func LoginHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	//Dependency Injection untuk mengambil uri param dan menyimpan dalam variabel
	nama := r.URL.Query().Get("nama")
	nim := r.URL.Query().Get("nim")

	//variabel yang menampung hasil cryptograph untuk mengamankan password
	hash := sha256.Sum256([]byte(nama + nim))

	// Mengambil cookie dengan nama Blokir dari browser user
	securityCookie, err := r.Cookie(listNamaCookie[1])

	// Cek jika cookie "blokir" tidak ada dan percobaan <= 3
	if err != nil && percobaan <= 3 {
		//variabel data yang menampung string return dari fungsi AksesDatabaseAdmin, dan err yang akan menampung error dari fungsi jika ada
		data, err := AksesDatabaseAdmin(db, nama, hex.EncodeToString(hash[:]))
		if err != nil {
			http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		if data == "" && percobaan <= 3 {
			data = "nama atau nim salah"
			percobaan++
			fmt.Fprintln(w, "nama atau nim salah\nsisa percobaan : ", strconv.Itoa(3-int(percobaan)))
		} else if percobaan < 3 {
			buatCookie(w, r, "login", "true")

			percobaan = 0
			log.Println("user berhasil masuk")
			// fmt.Fprintln(w, "selamat datang "+data)
			time.Sleep(3 * time.Second)
			http.Redirect(w, r, "/", http.StatusFound)
		}
	} else if err != nil && percobaan >= 3 {
		percobaan = 0
		// Jika percobaan login sudah 3 kali gagal, set cookie blokir
		buatCookie(w, r, "blokir", "true")
	}

	// Jika cookie Blokir bernilai "true", redirect ke halaman 403
	if securityCookie != nil && securityCookie.Value == "true" {
		log.Println("user terblokir")
		http.Redirect(w, r, "/403", http.StatusForbidden)
		return
	}
}

func GetCookie(w http.ResponseWriter, r *http.Request) {
	// Mencoba mengambil cookie berdasarkan nama
	namaCookie := r.URL.Query().Get("q")
	if namaCookie == "" {

	}

	cookie, err := r.Cookie(namaCookie)
	if err != nil {
		fmt.Fprintln(w, "tidak ditemukan")
		return
	}

	// Menampilkan nilai cookie jika ditemukan
	fmt.Fprintf(w, "Cookie ditemukan! Nama: %s, Nilai: %s, Expired: %s\n", cookie.Name, cookie.Value, cookie.Expires)
}

func HandlingDashboard(w http.ResponseWriter, _ *http.Request, arr *[]string, arr2 *[]string) {
	for i := 0; i < len(*arr); i++ {
		fmt.Fprintf(w, "nama: %s nim : %s\n", (*arr)[i], (*arr2)[i])
	}
}

func TambahMahasiswa(w http.ResponseWriter, r *http.Request, db *sql.DB) error {
	URIquery := r.URL.Query()
	nama := URIquery.Get("nama")
	nim := URIquery.Get("nim")

	if nama == "" || nim == "" {
		fmt.Fprintln(w, "nama atau nim harus diberikan")
	}

	Nim, err := strconv.Atoi(nim)

	if err != nil {
		return err
	}

	TambahUserDatabase(db, nama, uint(Nim))
	fmt.Fprintln(w, nama+"dengan nim: "+nim+"berhasil dimasukan")
	log.Println(nama + "dengan nim: " + nim + "berhasil dimasukan")
	return nil
}

func HapusMahasiswa(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	cekCookie, err := r.Cookie("login")
	if err != nil || cekCookie.Value == "false" {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	nim := r.URL.Query().Get("nim")
	nama := r.URL.Query().Get("nama")

	err = HapusUserDatabase(db, nama, nim)
	if err != nil {
		log.Fatal(err)
		return
	}
	if nim != "" && nama != "" {
		fmt.Fprintln(w, "data mahasiswa dengan nama: "+nama+"dengan nim: "+nim+"berhasil dihapus")
		log.Println("data mahasiswa dengan nama: " + nama + "dengan nim: " + nim + "berhasil dihapus")
	} else if nama != "" {
		fmt.Fprintln(w, "data mahasiswa dengan nama: "+nama+"berhasil dihapus")
		log.Println("data mahasiswa dengan nama: " + nama + "berhasil dihapus")
	} else if nim != "" {
		fmt.Fprintln(w, "data mahasiswa dengan nim: "+nim+"berhasil dihapus")
		log.Println("data mahasiswa dengan nim: " + nim + "berhasil dihapus")
	}
}

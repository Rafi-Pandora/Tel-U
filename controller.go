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

func cekCookie(_ http.ResponseWriter, r *http.Request) *[][]string {
	//deklarasi array multidimensi untuk menyimpan return value fungsi dan boolean
	var nilaiReturn = [][]string{}
	var adaCookie bool

	//loop bedasarkan nama cookie
	for i := 0; i < len(listNamaCookie); i++ {
		cekCookie, err := r.Cookie(listNamaCookie[i])

		//pengecekan cookie ada atau tidaknya
		if err != nil {
			nilaiReturn = append(nilaiReturn, []string{"", ""}) //jika tidak ada masukan nilai string kosong ke dimensi kedua array

			//percabangan jika cookie tersebut ada
		} else {
			//loop elemen bedasarkan value cookie
			for j := 0; j < len(listValueCookie[j]); j++ {
				/*percabangan untuk mengecek apakah value dari cookie sama dengan array listValueCookie,
				jika ada ubah nilai ada cookie ke true kemudian keluar dari looping*/
				if cekCookie.Value == listValueCookie[j] {
					adaCookie = true
					break
				}
			}
			if adaCookie {
				nilaiReturn = append(nilaiReturn, []string{listNamaCookie[i], cekCookie.Value})
			} else {
				nilaiReturn = append(nilaiReturn, []string{listNamaCookie[i], ""})
			}
		}
	}
	return &nilaiReturn
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

func LogOut(w http.ResponseWriter, r *http.Request) {
	cekCookie, err := r.Cookie("login")
	if err != nil || cekCookie.Value == "false" {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	cookie := &http.Cookie{
		Name:     "login",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add(-60 * time.Second),
	}
	http.SetCookie(w, cookie)
	// fmt.Fprintln(w, "mencoba logout...")
	time.Sleep(3 * time.Second)
	http.Redirect(w, r, "/get", http.StatusTemporaryRedirect)
}

func GetCookie(w http.ResponseWriter, r *http.Request) {
	// Mencoba mengambil cookie berdasarkan nama
	namaCookie := r.URL.Query().Get("q")
	if namaCookie == "" {
		//memanggil fungsi cek cookie dan memasukan nilainya kedalam variabel list
		listCookie := cekCookie(w, r)

		//looping array pointer multidimensi untuk mencetak nilai
		//looping dimensi pertama array
		for i := 0; i < len(*listCookie); i++ {
			//looping dimensi kedua array
			for j := 0; j < len((*listCookie)[i]); j++ {
				// Periksa apakah nilai cookie kosong
				if (*listCookie)[i][j] == "" {
					fmt.Fprintf(w, "Cookie tidak ditemukan!")
					return // keluar dari fungsi
				} else {
					// Menampilkan nama dan nilai cookie yang ditemukan
					fmt.Fprintf(w, "Cookie ditemukan! Nama: %s, Nilai: %s\n", (*listCookie)[i][0], (*listCookie)[i][1])
					return
				}
			}
		}
		return //keluar dari fungsi
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
	cekCookie, err := r.Cookie("login")
	if err != nil || cekCookie.Value == "false" {
		http.Redirect(w, r, "/login", http.StatusFound)
		return nil
	}
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

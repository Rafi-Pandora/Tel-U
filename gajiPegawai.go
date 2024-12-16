package main

import (
	"strconv"
	"strings"
)

func GajiPegawai(Jabatan string, MasaKerja uint, JumlahAnak uint) string {
	var gajiPokok uint
	var tunjangan uint = 0

	Jabatan = strings.ToLower(Jabatan)

	if JumlahAnak > 3 {
		JumlahAnak = 3
	}

	switch Jabatan {
	case "staf":
		if MasaKerja < 5 {
			gajiPokok = 4000
		} else if MasaKerja > 5 && MasaKerja < 10 {
			gajiPokok = 4000
			tunjangan = 100
		} else {
			gajiPokok = 5000
			tunjangan = 100
		}
	case "direktur":
		gajiPokok = 20000
		tunjangan = 500
	case "manager":
		tunjangan = 300
		if MasaKerja > 10 {
			gajiPokok = 10000
		} else {
			gajiPokok = 8500
		}
	default:
		GajiPegawai(InputStr("\n\ninput salah!, input yang diizinkan: Staf|Direktur|Manager (string), Masa Kerja (uint), Jumlah anak (uint)\n\ngaji pegawai\nJabatan :"), uint(InputInt("Masa Kerja :")), uint(InputInt("Jumlah Anak :")))
		return ""
	}

	return strconv.Itoa(int(gajiPokok)) + " + " + strconv.Itoa(int(JumlahAnak*tunjangan)) + " = " + strconv.Itoa(int(gajiPokok+(JumlahAnak*tunjangan)))
}

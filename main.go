package main

func main() {
	println(Ktp(uint(InputInt("ktp\nmasukan usia :")), InputBool("memiliki Kartu keluarga? True|False")))
	println(TigaBialangan(InputInt("\n\nTiga bilangan\nmasukan angka a :"), InputInt("masukan angka b :"), InputInt("masukan angka c :")))
	println(ManagerEPL())
	println(GajiPegawai(InputStr("\n\ngaji pegawai\nJabatan :"), uint(InputInt("Masa Kerja :")), uint(InputInt("Jumlah Anak :"))))
	println(Ojol(InputInt("\n\nOjol\nMasukan jam:"), InputInt("Masukan menit:"), float64(InputInt("Masukan Jarak:"))))
}

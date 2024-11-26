package main

import "strings"

func soal1(char string) string {
	var konsonanArray = [5]string{"a", "i", "u", "e", "o"}
	const konsonan string = "konsonan"
	const bukanKonsonan string = "bukan konsonan"

	if len(char) != 1 {
		return "hanya masukan 1 karakter saja"
	}

	for i := 0; i < len(konsonanArray); i++ {
		if char == konsonanArray[i] || char == strings.ToUpper(konsonanArray[i]) {
			return konsonan
		}
	}
	return bukanKonsonan
}

package main

import "fmt"

func cetakInt(pesan string) int {
	var i int
	println("\n\n", pesan)
	fmt.Scanln(&i)
	return i
}

func cetakStr(pesan string) string {
	var i string
	println("\n\n", pesan)
	fmt.Scanln(&i)
	return i
}

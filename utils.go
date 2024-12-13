// package atau class yang berisi fungsi pendukung
package main

import "fmt"

func InputInt(pesan string) int {
	var i int
	println(pesan)
	fmt.Scanln(&i)
	return i
}

func InputStr(pesan string) string {
	var i string
	println(pesan)
	fmt.Scanln(&i)
	return i
}

func InputBool(pesan string) bool {
	var i bool
	println(pesan)
	fmt.Scanln(&i)
	return i
}

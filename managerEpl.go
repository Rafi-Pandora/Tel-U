package main

import (
	"strconv"
	"strings"
)

func ManagerEPL() string {
	println("\n")

	var matchArr = [5]string{}
	var kalah uint = 0

	for i := 0; i < len(matchArr); i++ {
		matchArr[i] = strings.ToLower(InputStr("match ke :" + strconv.Itoa(i+1)))

		if matchArr[i] == "kalah" {
			kalah++
		}
	}

	if kalah == 5 {
		return "dipecat"
	}
	return "tidak dipecat"
}

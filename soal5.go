package main

import "sort"

func soal5(input0 int, input1 int, input2 int) bool {

	inputArr := []int{input0, input1, input2}
	sort.Ints(inputArr)
	entryNum, midNum, endNum := inputArr[0], inputArr[1], inputArr[2]

	numList := make([]int, endNum-entryNum)
	for i := range numList {
		numList[i] = i + entryNum
	}

	mid := len(numList) / 2

	if mid%2 != 0 && midNum == numList[mid] {
		return true
	} else {
		return false
	}
}

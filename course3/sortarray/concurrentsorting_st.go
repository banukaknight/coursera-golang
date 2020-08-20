package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	inputSli := input()
	divSli := divideArray(inputSli, 4)
	slices := make(chan []int)

	for i, v := range divSli {
		go func(sli []int, i int) {
			sortedSli := bubbleSort(sli)
			fmt.Printf("Goroutine №%d: %v\n", i+1, sortedSli)
			slices <- sortedSli
		}(v, i)
	}

	var notMergedSli [][]int

	for i := 0; i < len(divSli); i++ {
		notMergedSli = append(notMergedSli, <-slices)
	}

	result := unpackMerge(notMergedSli)

	fmt.Println(result)

}

func input() []int {

	var sli []int
	var stopFlag string

	for i := 0; ; i++ {
		fmt.Println("Input an integer or \"stop\" to stop inputing")
		for {
			var temp string
			fmt.Scan(&temp)
			elem, err := strconv.ParseInt(temp, 10, 64)
			if err != nil {
				if temp == "stop" || temp == "Stop" {
					stopFlag = "stop"
					break
				} else {
					fmt.Println("Incorrect data, please, try again")
				}
			} else {
				sli = append(sli, int(elem))
			}

		}
		if stopFlag == "stop" || stopFlag == "Stop" {
			break
		}
	}

	return sli
}

func divideArray(sli []int, n int) [][]int {
	newLen := int(math.Floor(float64(len(sli) / n)))
	rest := len(sli) % n
	var superSlice [][]int
	var counter int = 0
	for i := 0; i < n; i++ {
		var temp []int
		if i != n-1 {
			temp = sli[counter : counter+newLen]
			counter = counter + newLen
		} else {
			temp = sli[counter : counter+newLen+rest]

		}
		superSlice = append(superSlice, temp)
	}
	return superSlice
}

func bubbleSort(sli []int) []int {

	n := len(sli)

	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if sli[j] > sli[j+1] {
				sli[j], sli[j+1] = sli[j+1], sli[j]
			}
		}
	}

	return sli
}

func mergeSlices(sli [][]int) [][]int {
	var resSli [][]int
	if len(sli) > 1 {
		sli1, sli2 := sli[0], sli[1]
		resSli = sli[2:len(sli)]
		var mergedSli []int = make([]int, 0)
		i, j := 0, 0

		//Запись основной части элементов
		for i < len(sli1) && j < len(sli2) {
			if sli1[i] < sli2[j] {
				mergedSli = append(mergedSli, sli1[i])
				i++
			} else {
				mergedSli = append(mergedSli, sli2[j])
				j++
			}
		}

		//Запись остатка елементов в массив
		if i == len(sli1) {
			for j < len(sli2) {
				mergedSli = append(mergedSli, sli2[j])
				j++
			}
		} else if j == len(sli2) {
			for i < len(sli1) {
				mergedSli = append(mergedSli, sli1[i])
				i++
			}
		}

		resSli = append(resSli, mergedSli)

		if len(resSli) > 1 {

			resSli = mergeSlices(resSli)

		}

	}

	return resSli
}

func unpackMerge(sli [][]int) []int {
	resSli := mergeSlices(sli)

	return resSli[0]
}

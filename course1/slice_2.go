package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var numsSlice = make([]int, 3) // Slice of length 3 and capacity 3
	var input string
	var num int
	var i int = 0

	for {
		fmt.Printf("Please enter an integer or X to exit->")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input = scanner.Text()
		if input == "X" || input == "x" {
			fmt.Println("Thanks bye!")
			break
		}
		num, _ = strconv.Atoi(input)
		if i > 2 {
			numsSlice = append(numsSlice, num)
		} else {
			numsSlice[0] = num // always feed the new entry at 0th position
			i++
		}
		// sort all the numbers
		sort.Ints(numsSlice)
		fmt.Println("The sorted array in ascending order -> ", numsSlice)
	}
}

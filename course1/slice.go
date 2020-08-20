package main

import "fmt"
import "strings" //string func
import "strconv" //convert string to other
import "sort" //sort func

func main() {
	slice1 := make([]int, 0)
	var str1 string
	var int1 int
	var err error = nil
	fmt.Println("Enter an integer or x: ")
	
for true {
	err = nil //reset error
	
	_, err = fmt.Scanf("%s", &str1)
	
	if err != nil {
	//fmt.Println("String error: ", err)
	//"unexpected newline" error ignored
	}else if (strings.HasPrefix(str1, "x")){
		fmt.Println("Exit program")
		break
	}else{
		int1, err = strconv.Atoi(str1)
		if err == nil {
			//convert string to int and append
			slice1 = append(slice1, int1)
		}
		sort.Ints(slice1) //sort the slice
		fmt.Println(slice1) 
	}
	
}//end for loop
	
	fmt.Println("Final output is slice of length:", len(slice1) )
	fmt.Println(slice1) 
	
}//end main


/*Write a program which prompts the user to enter integers and stores the integers in a sorted slice. The program should be written as a loop. Before entering the loop, the program should create an empty integer slice of size (length) 3. 
During each pass through the loop, the program prompts the user to enter an integer to be added to the slice. The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order. The slice must grow in size to accommodate any number of integers which the user decides to enter. The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.*/
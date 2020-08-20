package main

import "fmt" //print & read func
import "strings" //string func
//import "bufio" //read line from user
//import "os" //file handle & read line
import "strconv" //atoi func

func BubbleSort( sli_ce []int){
	for i:=0; i<len(sli_ce)-1; i++{
		for j := 0; j< len(sli_ce)-1-i; j++ {			
			if ( sli_ce[j] > sli_ce[j+1] ){
				fmt.Println("swaps: ",sli_ce[i] , sli_ce[i+1])
				Swap( sli_ce, j) //call func to swap
			}
		}
	}
}//end BubbleSort

func Swap(sli_ce []int, i int) {
sli_ce[i] , sli_ce[i+1] = sli_ce[i+1] , sli_ce[i]
}//end swap

func main() {

	var st_num string
	var int_num int

	var s_len int
	mySlice := make([]int, 0)


	fmt.Println("Enter length of sequence (max:10): ")
	_,err := fmt.Scanf("%d", &s_len)

	if err != nil {
		fmt.Println("Error in length:", err)
	}else if(s_len >10)||(s_len==0){
		s_len = 10
	}
	
	fmt.Printf("Enter %d integers SEPERATED by SPACE: ",s_len)
	for i:=0; i<=s_len; i++{
	
		_, err := fmt.Scanf("%s", &st_num)
		
		int_num, err = strconv.Atoi(strings.TrimSpace(st_num))
		if err == nil {
			//convert string to int and append
			mySlice = append(mySlice, int_num)
		}	
	}
	
	fmt.Println(mySlice)
	
	BubbleSort(mySlice) //call function to sort
	
	fmt.Println("Sorted list is: ",mySlice)
	
	



}//end main

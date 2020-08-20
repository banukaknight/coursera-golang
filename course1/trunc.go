package main

import (
		"fmt" 		
)

func main() {

	//Scanf for scan string fromatting
	//Scanln for scan till end of line
	//scan functions store pointer variable

	var floatval float32

	fmt.Println("Enter a float value : ")
	_, err := fmt.Scanf("%f", &floatval)

	if err != nil {
	fmt.Println(err)
	}
	
	//type conversion
	intval := int(floatval)

	fmt.Printf("You have entered : %f \n", floatval)
	fmt.Printf("After Truncation : %d \n", intval)
	
	
	var myInt int
	fmt.Println("Input a floating point number: ")
	fmt.Scan(&myInt)
	fmt.Printf("%v, %T", myInt, myInt)

}
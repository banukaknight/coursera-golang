package main

import "fmt"
//import "strconv"
//import "strings"
import "bufio"
import "os"
import "encoding/json"


func main() {
  
	//var myMap map[string]string
	//myMap = make(map[string]string) //create empty map

	var name string
	var address string

	fmt.Printf("Please enter a name ->")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name = scanner.Text()
	fmt.Printf("Enter address ->")
	scanner.Scan()
	address = scanner.Text()

	//create map in one line
	myMap := map[string]string {"name":name, "address":address}

	fmt.Println("Map variable ->", myMap)
	
	jsonarr,_ := json.Marshal(myMap)
	
	fmt.Println("\nRAW JSON OBJECT ->\n", jsonarr)
	
	fmt.Println("\nString JSON OBJECT ->\n")
	os.Stdout.Write(jsonarr)
}

/*Write a program which prompts the user to first enter a name, and then enter an address. Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively. Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.

Submit your source code for the program, “makejson.go”.*/
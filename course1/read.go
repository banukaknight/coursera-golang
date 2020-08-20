package main

import (
	"bufio"
	"fmt"
	"os"
	//"sort"
	"strings"
	//"strconv"
	//"github.com/astaxie/beego/validation"
	"log" //log stuff
	//"text/tabwriter"
)

type nameStrc struct {
	fname string `valid:MaxSize(20)`
	lname string `valid:MaxSize(20)`
} 

func main() {

	mySlice := make([]nameStrc, 0)
	//fname := "myfile.txt"
	fmt.Printf("Please enter file name -> (expect: myfile.txt) :")
	scan1 := bufio.NewScanner(os.Stdin)
	scan1.Scan()
	fname := scan1.Text()
		
	file, err := os.Open(fname)
    if err != nil {
        log.Fatal(err)
    }
	defer file.Close()
	
	scan2 := bufio.NewScanner(file)
    for scan2.Scan() {
        thisline := scan2.Text()
		fmt.Println("Read line -> ", thisline)
		
		trimedline := strings.TrimSpace(thisline)
		sliceline := strings.Split(trimedline, " ")
		//fmt.Println(sliceline)
		
		//ignore name without atleast 2 name parts
		if len(sliceline) >= 2 {
		strcObj := nameStrc{sliceline[0], sliceline[1]}
		mySlice = append(mySlice, strcObj)
		}
    }

    if err := scan2.Err(); err != nil {
        log.Fatal(err)
    }
	fmt.Println("-------------------------------------")
	fmt.Printf("%s \t\t| %s\n","First Name", "Last Name")
	fmt.Println("-------------------------------------")
	for _, x := range mySlice {
		//fmt.Println(x.fname, x.lname)
		
		fmt.Printf("%s \t\t\t| %s\n",x.fname, x.lname)
	}
	
	
}	
	

/*
Write a program which reads information from a file and represents it in a slice of structs. Assume that there is a text file which contains a series of names. Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.

Your program will define a name struct which has two fields, fname for the first name, and lname for the last name. Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file. Your program will successively read each line of the text file and create a struct which contains the first and last names found in the file. Each struct created will be added to a slice, and after all lines have been read from the file, your program will have a slice containing one struct for each line in the file. After reading all lines from the file, your program should iterate through your slice of structs and print the first and last names found in each struct.
*/
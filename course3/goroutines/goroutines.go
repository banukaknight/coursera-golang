package main

import (
	//"bufio"
	"fmt"
	//"os"
	//"sort"
	//"strings"
	//"strconv"
	//"github.com/astaxie/beego/validation"
	//"log" //log stuff
	//"text/tabwriter"
)

var x int = 5

func add2() {
    x = x+2
    fmt.Printf("\nadd1: %d",x)
}

func sub1() {
    x = x-1
    fmt.Printf("\nsub1: %d",x)
}

func main() {
    go add2()
    go sub1()
    
    fmt.Printf("\nfinx: %d",x)
    

    
}//end main
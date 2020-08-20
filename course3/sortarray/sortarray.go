package main

import (
	"fmt"
	"os"
    "sync"
)

var slice1 = make([]int, 0)
var slice2 = make([]int, 0)
var slice3 = make([]int, 0)
var slice4 = make([]int, 0)

var slice5 = make([]int, 0)
var slice6 = make([]int, 0)
var slice7 = make([]int, 0)

var i = 0
var j = 0

//function to sort a slice
func sort_slice(sli_ce []int, wg *sync.WaitGroup){
    for i=0;i<len(sli_ce); i=i+1{
        for j=i+1;j<len(sli_ce); j=j+1{
            if sli_ce[i]>sli_ce[j]{
                sli_ce[i],sli_ce[j] = sli_ce[j],sli_ce[i]
            }
        }
    }
    fmt.Println("Go routine post-sort:", sli_ce)
    wg.Done()
}

//function to merge 2 slices into one
func merge_slice(sl_1 []int, sl_2 []int, slice_out *[]int){
    i,j = 0,0
    for (i<len(sl_1) && j<len(sl_2)){
        if sl_1[i]<sl_2[j] {
            *slice_out = append(*slice_out, sl_1[i])
            i++
        }else{
            *slice_out = append(*slice_out, sl_2[j])
            j++
        }
    }    
    for i<len(sl_1){
        *slice_out = append(*slice_out, sl_1[i])
        i++
    }
    for j<len(sl_2){
        *slice_out = append(*slice_out, sl_2[j])
        j++
    }
}


func main() {
    var err error
    var arrlen int
    var elem int
    
	fmt.Println("Enter length of array: ")
	_, err = fmt.Scan(&arrlen)
    if err != nil {
	   fmt.Println(err)
	}
    if arrlen < 4{
        fmt.Println("Need a list with 4+ elements")
        os.Exit(3)
    }
    
    var arr123len int
    var arr4len int
    arr123len = arrlen/4
    arr4len = arrlen-(arr123len*3)
    
    fmt.Println("Enter", arrlen, "integers in a single line seperated by spaces: ")
    for i:=0;i<arr123len;i++{
        _, err = fmt.Scan(&elem)
        if err == nil {
            slice1 = append(slice1, elem)
        }else{
            i--
        }
    }
    for i:=0;i<arr123len;i++{
        _, err = fmt.Scan(&elem)
        if err == nil {
            slice2 = append(slice2, elem)
        }else{
            i--
        }
    }
    for i:=0;i<arr123len;i++{
        _, err = fmt.Scan(&elem)
        if err == nil {
            slice3 = append(slice3, elem)
        }else{
            i--
        }
    }
    for i:=0;i<arr4len;i++{
        _, err = fmt.Scan(&elem)
        if err == nil {
            slice4 = append(slice4, elem)
        }else{
            i--
        }
    }
    
    //wait group to keep counter
    var wg sync.WaitGroup

    //sort 4 slices
    fmt.Println("Pre-Sorting 4x arrays as inserted by user: ")
    fmt.Println(slice1, slice2, slice3, slice4,"\n")
    wg.Add(1)
    go sort_slice(slice1,&wg)
    wg.Add(1)
    go sort_slice(slice2,&wg)
    wg.Add(1)
    go sort_slice(slice3,&wg)
    wg.Add(1)
    go sort_slice(slice4,&wg)
    
    wg.Wait()
    //only after the 4 sorting goroutines are done program will continue
    fmt.Println("\nPost-Sorting 4x sub arrays: ")
    fmt.Println(slice1, slice2, slice3, slice4)
    
    //then merge 2 x 2 then into one
    merge_slice(slice1, slice2, &slice5)
    merge_slice(slice3, slice4, &slice6)
    //final merge
    merge_slice(slice5, slice6, &slice7)
 
    
    fmt.Println("Merged 2x sub arrays: ")
    fmt.Println(slice5,slice6)
    fmt.Println("\nFinal merged single sorted array: ")
    fmt.Println(slice7)
    
    
    fmt.Println("\n\t\tProgram by: Banuka Vidusanka")
}//end main
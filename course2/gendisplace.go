//by Banuka Vidusanka //

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

//func GenDisplaceFn accept (acceleration, velocity, displacement)
func GenDisplaceFn(a float64, v float64, s float64) func (float64) float64 {
    
    fn := func (t float64) float64{
        return (0.5*a*t*t + v*t + s)        
    }
    return fn
}//end GenDisplaceFn

func main() {
    var acce float64 //acceleration
    var velo float64 //initial velocity
    var sdis float64 //initial displacement
    var time float64 //time
    
    var err error

	fmt.Println("Enter acceleration: ")
	_, err = fmt.Scan(&acce)
    if err != nil {
	   fmt.Println(err)
	}
    
    fmt.Println("Enter initial velocity: ")
	_, err = fmt.Scan(&velo)
    if err != nil {
	   fmt.Println(err)
	}
    
    fmt.Println("Enter initial displacement: ")
	_, err = fmt.Scan(&sdis)
    if err != nil {
	   fmt.Println(err)
	}
    
    
    //displacefunciton that accept time as input
    disfn := GenDisplaceFn(acce, velo, sdis)
    
    
    fmt.Println("Enter time: ")
	_, err = fmt.Scan(&time)
    if err != nil {
	   fmt.Println(err)
	}
    
    fmt.Println(disfn(time))
    
    
}//end main


/*Let us assume the following formula for displacement s as a function of time t, acceleration a, initial velocity vo, and initial displacement so.

s =½ a t2 + vot + so

Write a program which first prompts the user to enter values for acceleration, initial velocity, and initial displacement. Then the program should prompt the user to enter a value for time and the program should compute the displacement after the entered time.

You will need to define and use a function called GenDisplaceFn() which takes three float64 arguments, acceleration a, initial velocity vo, and initial displacement so. GenDisplaceFn() should return a function which computes displacement as a function of time, assuming the given values acceleration, initial velocity, and initial displacement. The function returned by GenDisplaceFn() should take one float64 argument t, representing time, and return one float64 argument which is the displacement travelled after time t.

For example, let’s say that I want to assume the following values for acceleration, initial velocity, and initial displacement: a = 10, vo = 2, so = 1. I can use the following statement to call GenDisplaceFn() to generate a function fn which will compute displacement as a function of time.

fn := GenDisplaceFn(10, 2, 1)

Then I can use the following statement to print the displacement after 3 seconds.

fmt.Println(fn(3))

And I can use the following statement to print the displacement after 5 seconds.

fmt.Println(fn(5))
*/
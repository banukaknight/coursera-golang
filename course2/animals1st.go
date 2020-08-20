//by Banuka Vidusanka //

package main
import (
	"bufio"
	"fmt"
	"os"
	//"sort"
	"strings"
	//"strconv"
	//"github.com/astaxie/beego/validation"
	//"log" //log stuff
	//"text/tabwriter"
)

type Animal struct {
    food string
    locomotion string
    noise string
}//end struct

func (ani *Animal) Eat(){
    fmt.Printf(" likes to eat: %s\n", ani.food)
}

func (ani *Animal) Move(){
    fmt.Printf("'s locomotion method is: %s\n", ani.locomotion)
}

func (ani *Animal) Speak(){
    fmt.Printf(" makes the noise: %s\n", ani.noise)
}

//function that takes obj pointer as argument
//calls another func
//no return type
func (ani *Animal) callAction(act string) {
    switch act {
        case "eat":
            {
                ani.Eat()
            }
        case "move":
            {
                ani.Move()
            }
        case "speak":
            {
                ani.Speak()
            }
        default:
            {
                fmt.Printf(" doesnt have action: %s",act)
            }
    }//end switch
}//end callAction func

func main() {
    fmt.Println("------------------------------")
    
    //objects initialization
    cow := Animal{"grass","walk","moo"}
    bird := Animal{"worms","fly","peep"}
    snake := Animal{"mice","slither","hsss"}
    
    //map (hashtable) for matching string to obj
    aniMap := map[string] Animal {
        "cow":cow,
        "bird":bird,
        "snake":snake,
    }   
    
    scan1 := bufio.NewScanner(os.Stdin)
    fmt.Printf("\n Enter NAME and ACTION of Animal seperated by space\n")
    fmt.Printf(" Options: (cow,bird,snake) (eat,move,speak) \n Enter x to exit program -> ")
    
    for {
        scan1.Scan()
        u_input := scan1.Text()

        trimedline := strings.TrimSpace(u_input)
        lower_line := strings.ToLower(trimedline)
        slice_line := strings.Split(lower_line, " ")

        if(slice_line[0]=="x"){
            //os.Exit(3)
            break
        }else if(len(slice_line)<2){
            fmt.Printf("Not enough arguments\n")
            fmt.Printf("\n -> ")
            continue
        }
        //fmt.Println(slice_line[0] , slice_line[1])

        
        //match input animal to obj from map
        o_Ani,ok := aniMap[slice_line[0]]
        if ok == false {
            fmt.Printf("Animal NAME not found\n")
            fmt.Printf(" -> ")
            continue
        }

        //call action func with o_ani as argument
        fmt.Printf("\n\t%s",slice_line[0])
        o_Ani.callAction(slice_line[1])
        fmt.Printf(" -> ")
    
    }//end forever for loop
    
    fmt.Println("--------------made by Banuka----------------")
    fmt.Printf("End of Program\n")
}//end main

/*
Write a program which allows the user to get information about a predefined set of animals. Three animals are predefined, cow, bird, and snake. Each animal can eat, move, and speak. The user can issue a request to find out one of three things about an animal: 1) the food that it eats, 2) its method of locomotion, and 3) the sound it makes when it speaks. The following table contains the three animals and their associated data which should be hard-coded into your program.

Animal	Food eaten	Locomotion method	Spoken sound
cow	grass	walk	moo
bird	worms	fly	peep
snake	mice	slither	hsss
Your program should present the user with a prompt, “>”, to indicate that the user can type a request. Your program accepts one request at a time from the user, prints out the answer to the request, and prints out a new prompt. Your program should continue in this loop forever. Every request from the user must be a single line containing 2 strings. The first string is the name of an animal, either “cow”, “bird”, or “snake”. The second string is the name of the information requested about the animal, either “eat”, “move”, or “speak”. Your program should process each request by printing out the requested data.
You will need a data structure to hold the information about each animal. Make a type called Animal which is a struct containing three fields:food, locomotion, and noise, all of which are strings. Make three methods called Eat(), Move(), and Speak(). The receiver type of all of your methods should be your Animal type. The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion, and the Speak() method should print the animal’s spoken sound. Your program should call the appropriate method when the user makes a request.
*/    
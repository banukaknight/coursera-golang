//by Banuka Vidusanka DATE: 2020-08-15

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

type Animal interface {
    Eat()
    Move()
    Speak()
}//end interface

type Cow struct {food, locomotion, noise string}
func (a Cow) Eat() {fmt.Printf("the Cow likes to eat: %s\n", a.food)}
func (a Cow) Move() {fmt.Printf("the Cow's locomotion method is: %s\n", a.locomotion)}
func (a Cow) Speak() {fmt.Printf("the Cow makes the noise: %s\n", a.noise)}

type Bird struct {food, locomotion, noise string}
func (a Bird) Eat() {fmt.Printf("the Bird likes to eat: %s\n", a.food)}
func (a Bird) Move() {fmt.Printf("the Bird's locomotion method is: %s\n", a.locomotion)}
func (a Bird) Speak() {fmt.Printf("the Bird makes the noise: %s\n", a.noise)}

type Snake struct {food, locomotion, noise string}
func (a Snake) Eat() {fmt.Printf("the Snake likes to eat: %s\n", a.food)}
func (a Snake) Move() {fmt.Printf("the Snake's locomotion method is: %s\n", a.locomotion)}
func (a Snake) Speak() {fmt.Printf("the Snake makes the noise: %s\n", a.noise)}


func getAniInfo(ani Animal, act string){
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
                fmt.Printf("Info: %s is not recognized.\nExpecting: eat,move,speak\n",act)
            }
    }//end switch
    
}

func main() {
    fmt.Println("------------------------------------------")
    fmt.Println("------------Welcome to My Barn------------")
    //copy of these 3 objects will be used & stored in anilist slice for each new animal created
    Cow_0 := Cow{food:"grass", locomotion:"walk", noise:"moo"}
    Bird_0 := Bird{food:"worms", locomotion:"fly", noise:"peep"}
    Snake_0 := Snake{food:"mice", locomotion:"slither", noise:"hsss"}
    
    //anilist is a slice for holding each new animal cow,bird,snake added
    anilist := make([]Animal, 0)
    //nameMap help map given animal name to created animal object
    nameMap := make (map[string]Animal)
    //scan1 scanner will scan user input and process
    scan1 := bufio.NewScanner(os.Stdin)


    fmt.Printf("TO add a new Animal to System Enter : \n\t\t'newanimal' new_name_foranimal type_of_animal('cow,bird,snake')\n")
    fmt.Printf("\tEx: newanimal mybirdyy bird\n")
    fmt.Printf("TO query about existing animal Enter: \n\t\t'query' name_of_animal info_expected('move,eat,speak')\n")
    fmt.Printf("\tEx: query mybirdyy eat\n")
    fmt.Printf("\t\tTo exit program Enter: 'x'\n")
    fmt.Println("------------------------------------------")
    fmt.Printf("\nExpect Input -> ")

    //forloop will run until input = 'x'
    for scan1.Scan(){
        u_input := scan1.Text()
        trimedline := strings.TrimSpace(u_input)
        lower_line := strings.ToLower(trimedline)
        slice_line := strings.Split(lower_line, " ")
        //user input is trimed, lowercased & then splitted to an array-slice of strings (expecting either 1 string = 'x' or three strings with commands)
        
        
        if(slice_line[0]=="x"){
            //os.Exit(3)
            break
        }else if(len(slice_line)<3){
            fmt.Printf("Not enough arguments\n")
            fmt.Printf("\n -> ")
            continue
        }
        
        
        if(slice_line[0]=="newanimal"){
            //slice_line[0]->"newanimal"
            //slice_line[1]->new_name_for_animal 
            //slice_line[2]->"cow/bird/snake"
            if(slice_line[2]=="cow"){
                anilist = append(anilist, Cow_0)
                nameMap[slice_line[1]] = anilist[len(anilist)-1]
            }else if(slice_line[2]=="bird"){
                anilist = append(anilist, Bird_0)
                nameMap[slice_line[1]] = anilist[len(anilist)-1]
            }else if(slice_line[2]=="snake"){
                anilist = append(anilist, Snake_0)
                nameMap[slice_line[1]] = anilist[len(anilist)-1]
            }else{
                fmt.Printf("Animal not recognized.\nExpecting: cow,bird,snake\n")
                fmt.Printf("\n -> ")
                continue
            }
            fmt.Printf("%s the new %s added Successfuly to System.\n", slice_line[1], slice_line[2])
            fmt.Printf("Created it!\n")
            
        }else if(slice_line[0]=="query"){
            //slice_line[0]->"query"
            //slice_line[1]->seeking_animal_name 
            //slice_line[2]->"eat/move/speak" 
            if nameMap[slice_line[1]] == nil {
                //check if animal with given name is found in nameMap MAP
                fmt.Printf("Animal named: %s not found.\n",slice_line[1])
                fmt.Printf("\n -> ")
                continue
            }else{
                fmt.Printf("%s ",slice_line[1])
                getAniInfo(nameMap[slice_line[1]],slice_line[2])
            }
            
        }else{
            fmt.Printf("Wrong input pattern.\nExpecting: newanimal or query as first string\n")
            fmt.Printf("\n -> ")
            continue
        }
    
        fmt.Printf("\n -> ")
    }//end for-loop
        
        
    fmt.Println("--------------made by Banuka----------------")
    fmt.Println("----------instagram.com/theBKrox----------------")
    fmt.Printf("End of Program\n")
}//end main


//TOOK ME 2 days to finish this program.
//hope it is upto standards. THANK YOU FOR REVIEWING my CODE
//instagram: instagram.com/theBKrox

/*
Write a program which allows the user to create a set of animals and to get information about those animals. Each animal has a name and can be either a cow, bird, or snake. With each command, the user can either create a new animal of one of the three types, or the user can request information about an animal that he/she has already created. Each animal has a unique name, defined by the user. Note that the user can define animals of a chosen type, but the types of animals are restricted to either cow, bird, or snake. The following table contains the three types of animals and their associated data.

Animal	Food eaten	Locomtion method	Spoken sound
cow	grass	walk	moo
bird	worms	fly	peep
snake	mice	slither	hsss
Your program should present the user with a prompt, “>”, to indicate that the user can type a request. Your program should accept one command at a time from the user, print out a response, and print out a new prompt on a new line. Your program should continue in this loop forever. Every command from the user must be either a “newanimal” command or a “query” command.

Each “newanimal” command must be a single line containing three strings. The first string is “newanimal”. The second string is an arbitrary string which will be the name of the new animal. The third string is the type of the new animal, either “cow”, “bird”, or “snake”. Your program should process each newanimal command by creating the new animal and printing “Created it!” on the screen.

Each “query” command must be a single line containing 3 strings. The first string is “query”. The second string is the name of the animal. The third string is the name of the information requested about the animal, either “eat”, “move”, or “speak”. Your program should process each query command by printing out the requested data.

Define an interface type called Animal which describes the methods of an animal. Specifically, the Animal interface should contain the methods Eat(), Move(), and Speak(), which take no arguments and return no values. The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion, and the Speak() method should print the animal’s spoken sound. Define three types Cow, Bird, and Snake. For each of these three types, define methods Eat(), Move(), and Speak() so that the types Cow, Bird, and Snake all satisfy the Animal interface. When the user creates an animal, create an object of the appropriate type. Your program should call the appropriate method when the user issues a query command.
*/    
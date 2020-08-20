package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Animal data type
type Animal struct {
	food       string
	locomotive string
	noise      string
}

// Eat pritns food
func (a *Animal) Eat() {
	fmt.Println(a.food)
}

// Move prints movement
func (a *Animal) Move() {
	fmt.Println(a.locomotive)
}

// Speak prints noise
func (a *Animal) Speak() {
	fmt.Println(a.noise)
}

/* 1. Using map data structure to hardcode respective animal properties */
func data(animal string) Animal {
	animalData := map[string]Animal{
		"cow": {
			food:       "grass",
			locomotive: "walk",
			noise:      "moo",
		},
		"bird": {
			food:       "worms",
			locomotive: "fly",
			noise:      "peep",
		},
		"snake": {
			food:       "mice",
			locomotive: "slither",
			noise:      "hsss",
		},
	}

	// returns data type of Animal
	return animalData[animal]
}

/* 2. function to execute user picked action against animal type */
func action(act string, a *Animal) {
	actMap := map[string]func(){
		"move":  a.Move,
		"eat":   a.Eat,
		"speak": a.Speak,
	}

	if actMap[act] == nil {
		fmt.Println(act, "action not supported")
		return
	}

	// now we have respective function and will execute it
	actMap[act]()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	for scanner.Scan() {
		inputString := scanner.Text()

		if inputString == "x" {
			fmt.Println("exiting...")
			break
		}

		// 0.1 Splitting by space to seperate animal & action
		inputArr := strings.Split(inputString, " ")

		// 0.2 Building animal a of user entered
		pickedAnimal := data(inputArr[0]) // data helper function used
		a := Animal{
			food:       pickedAnimal.food,
			locomotive: pickedAnimal.locomotive,
			noise:      pickedAnimal.noise,
		}

		// 0.3 Executing user entered action with generated animal a
		pickedAction := inputArr[1]
		action(pickedAction, &a) // action helper function used

		// 0.4 Re-Prompting user for to enter data
		fmt.Print("> ")
	}
}

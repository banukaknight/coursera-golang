package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var animals = make(map[string]*Animal)

type Animal struct {
	food       string
	locomotion string
	sound      string
}

func NewAnimal(food string, locomotion string, sound string) *Animal {
	return &Animal{food: food, locomotion: locomotion, sound: sound}
}

func (animal *Animal) eat() string {
	return animal.food
}

func (animal *Animal) move() string {
	return animal.locomotion
}

func (animal *Animal) speak() string {
	return animal.sound
}

func (animal *Animal) makesAction(action string) string {
	switch action {

	case "eat":
		{
			return animal.eat()
		}

	case "move":
		{
			return animal.move()
		}
	case "speak":
		{
			return animal.speak()
		}
	default:
		resp := fmt.Sprintf("There is no action `%s`", action)
		return resp
	}
}

func initAnimals() {
	cow := NewAnimal("grass", "walk", "moo")
	bird := NewAnimal("worms", "fly", "peep")
	snake := NewAnimal("mice", "slither", "hsss")
	animals["cow"] = cow
	animals["bird"] = bird
	animals["snake"] = snake
}

func main() {
	initAnimals()
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("DIGIT COMMAND (ex. cow move) or 'q' to exit \n > ")
		scanner.Scan()

		//clean extra spaces
		input := strings.Trim(scanner.Text(), " ")
		matcher := regexp.MustCompile(`\s{2,}`)
		input = matcher.ReplaceAllLiteralString(input, " ")

		words := strings.Split(input, " ")
		if len(words) < 2 {
			fmt.Print("INVALID COMMAND!. command should have two words, ex `cow eat` \n\n")
			continue
		}
		animalName := strings.ToLower(words[0])
		action := strings.ToLower(words[1])

		animal := animals[animalName]

		if animal == nil {
			fmt.Printf("animal `%s` wasn't found in out set. Try wit another animal! \n\n", animalName)
			continue
		}

		fmt.Println("-------------------- ACTION --------------------")
		fmt.Println(animal.makesAction(action))
		fmt.Print("-------------------- END --------------------\n\n")

	}

}

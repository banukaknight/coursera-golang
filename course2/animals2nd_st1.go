package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
)

type Animal interface{
  Eat()
  Move()
  Speak()
}

type Sample struct {
  food,locomotion,speech string
}

func (smp Sample) Eat() {
    fmt.Println(smp.food)
}

func (smp Sample) Move(){
   fmt.Println(smp.locomotion)
}

func (smp Sample) Speak() {
  fmt.Println(smp.speech)
}

func main()  {
    fmt.Println("Enter either a newanimal or query command")
    fmt.Println("For new animal command \"newAnimal <name for the animal> typeOfAnimal(cow or bird or snake)\"")
    fmt.Println("For query command \"query <name of the animal> tyoeOfAction(eat or speak or locomotion)\"")
    fmt.Println("Note : Follow the command instructions above carefully!!!!!")
    mp := make(map[string]Animal)
    cow := Sample{"grass","walk","moo"}
    bird := Sample{"worms","fly","peep"}
    snake := Sample{"mice","slither","hiss"}
    ref := make(map[string]Sample)
    ref["cow"] = cow
    ref["bird"] = bird
    ref["snake"] = snake
    for{
      fmt.Print("> ")
      scanner := bufio.NewScanner(os.Stdin)
      scanner.Scan()
      text := scanner.Text()
      arr := strings.Split(text," ")
      if arr[0] == "newAnimal"{
        mp[arr[1]] = ref[arr[2]]
        fmt.Println(arr[2] + " is created by the name "+arr[1])
      }else{
        tmp := mp[arr[1]]
        if arr[2] == "eat"{
          tmp.Eat()
        }else if arr[2] == "speak"{
          tmp.Speak()
        }else{
          tmp.Move()
        }
      }
    }
}

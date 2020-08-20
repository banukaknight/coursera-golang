package main

import "fmt" //print & read func
import "strings" //string func
import "bufio" //read line from user
import "os" //file handle & read line

func main() {

	//Scanf for scan string fromatting
	//Scanln for scan till end of line
	//scan functions store pointer variable
	
	reader := bufio.NewReader(os.Stdin)

	var myInput string

	fmt.Println("Enter a string : ")
	myInput,err := reader.ReadString('\n')

	if err != nil {
	fmt.Println(err)
	}
	
	//not using. too complicated
	//read line's last character will be newline. so -2
	//lastchar := string(lowInput[(len(lowInput))-3])
	
	inpStr := strings.TrimSpace(myInput)
	lowInput := strings.ToLower(inpStr)
	fmt.Printf("Checking String:\n"+ lowInput+"\n")
	
	
	switch {
		case !(strings.HasPrefix(lowInput, "i")):
			fmt.Printf("Doesn't start with: i\n")
			fmt.Printf("\nNot Found!\n")
		
		case !(strings.HasSuffix(lowInput, "n")):
			fmt.Printf("Last char is not: n\n")
			fmt.Printf("\nNot Found!\n")
			
		case !(strings.Contains(lowInput, "a")):
			fmt.Printf("String has no char: a\n")
			fmt.Printf("\nNot Found!\n")
			
		default:
			fmt.Printf("\nFound!\n")
	}

}
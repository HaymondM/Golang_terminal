package main

// https://gobyexample.com/range
import (
	"fmt"
	"math/rand"
	"os/user"
	"time"
)

func main() {

	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Println("Welcome " + user.Name)

	fmt.Print("If you need help do --options to see your options \n")
	for {
		fmt.Print("<< ")
		var input string
		fmt.Scanln(&input)
		switch input {
		case "--options":
			fmt.Print("Here are the options")
		case "rand":
			var x = (rand.Intn(77))          //1 way to declare a var
			rand.Seed(time.Now().UnixNano()) //to make a new rand num every time. Default seed is 1
			fmt.Println("This is a random number:", x)
		case "quit":
			fmt.Print("Bye")
			break
		default:
			fmt.Println("That is not a option")
		}
		//if strings.Compare(input, "options") == 1 {
		//	fmt.Print("Here are the options")
		//	break
		//}
		if input == "quit" {
			break
		}
	}
}

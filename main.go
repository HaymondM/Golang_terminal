package main

// https://gobyexample.com/range
import (
	"fmt"
	"math/rand"
	"net"
	"os"
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
		fmt.Scan(&input)
		switch input {
		case "--options":
			fmt.Println("Here are the options: ")
			fmt.Println("rand: get a random number")
			fmt.Println("Ping: ping a website to see if you get a request back")
			fmt.Println("Math: do some math")
			fmt.Println("Database: add remove and edit items in your database")
			fmt.Println("quit: end the terminal")
		case "rand":
			var x = (rand.Intn(77))          //1 way to declare a var
			rand.Seed(time.Now().UnixNano()) //to make a new rand num every time. Default seed is 1
			fmt.Println("This is a random number:", x)
		case "math":

		case "database":

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

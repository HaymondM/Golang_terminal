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

	fmt.Println("Welcome back " + user.Name)

	fmt.Printf("If you need help do --options to see your options \n")
	var input string
	fmt.Scan("<<", input)
	fmt.Printf("%s", input)

	var x = (rand.Intn(39)) //1 way to declare a var
	y := "plus"             //2nd way to declare a var

	rand.Seed(time.Now().UnixNano())           //to make a new rand num every time. Default seed is 1
	fmt.Println("This is a random number:", x) //adds a space by itself
	fmt.Println(9, y, 9, "=", 9+9)

}

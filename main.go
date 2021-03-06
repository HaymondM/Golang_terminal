package main

// https://gobyexample.com/range
import (
	"fmt"
	"math/rand"
	"os/user"
	"time"
)

//math choice
func plus(a float32, b float32) float32 {

	return a + b
}

func minus(a float32, b float32) float32 {

	return a - b
}

func multiply(a float32, b float32) float32 {

	return a * b
}

func divide(a float32, b float32) float32 {

	return a / b
}

//end math choice

//database
type person struct {
	name  string
	age   int
	phone string
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	//fmt.Println("Added!")
	return &p
}

func main() {

	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Println("Welcome " + user.Name)

	fmt.Print("If you need help type --options to see your options \n")
	for { //while loop
		fmt.Print("<< ")
		var input string
		fmt.Scan(&input)
		switch input {
		case "--options":
			fmt.Println("Here are the options: ")
			fmt.Println("rand: get a random number")
			fmt.Println("Math: do some math")
			fmt.Println("Database: add remove and edit items in your contact database")
			fmt.Println("quit: end the terminal")
		case "rand":
			var x = (rand.Intn(77))          //1 way to declare a var
			rand.Seed(time.Now().UnixNano()) //to make a new rand num every time. Default seed is 1
			fmt.Println("This is a random number:", x)
		case "math":
			var input_math string
			var num1 float32
			var num2 float32
			fmt.Println("Enter in your first number")
			fmt.Scan(&num1)
			fmt.Println("Enter in your second number")
			fmt.Scan(&num2)

			fmt.Println("What math do you want to do? +,-,*,/(type the sign)")
			fmt.Scan(&input_math)
			if input_math == "+" {
				res := plus(num1, num2)
				fmt.Println(num1, "+", num2, "=", res)
			} else if input_math == "-" {
				res := minus(num1, num2)
				fmt.Println(num1, "-", num2, "=", res)
			} else if input_math == "*" {
				res := multiply(num1, num2)
				fmt.Println(num1, "*", num2, "=", res)
			} else if input_math == "/" {
				res := divide(num1, num2)
				fmt.Println(num1, "/", num2, "=", res)

			} else {
				fmt.Println("You did not enter in the right symbol")
			}

		case "database":
			for { //while loop
				var a [10]person
				var name string
				var age int
				var phone string
				var input_data string
				s := person{name: name, age: age, phone: phone}
				count := 0
				fmt.Println(a)
				fmt.Println("Welcome to your contact database! Here are your options")
				fmt.Println("view: View the database")
				fmt.Println("remove: Remove data from the database")
				fmt.Println("add: Add to the database")
				fmt.Println("leave: Leave the database")
				fmt.Print("<< ")


				if input_data == "view" {
					fmt.Println("Here is the database:")
					fmt.Println(a)

				} else if input_data == "remove" {
					//cant get the rest of the data base finsihed
				} else if input_data == "add" {

					fmt.Println("What is the name of the person you want to add?")
					fmt.Scan((&name))
					fmt.Println("What is the age of the person?")
					fmt.Scan((&age))
					fmt.Println("What is the phone number of the person?(no spaces)")
					fmt.Scan(&phone)

					//wont work the vars seem to be local but not global
					s = person{name: name, age: age, phone: phone}
					a[0] = s
					count += 1

				} else if input_data == "leave" {
					fmt.Println("Goodbye")
					break
				}
			}
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
		if input == "quit" { //for the break of the while loop
			break
		}
	}
}

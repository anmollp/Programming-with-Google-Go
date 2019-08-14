package main

import (
"fmt"
"bufio"
"os"
"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}

func (c Cow) Eat() {
	fmt.Println("grass")
}

func (c Cow) Move() {
	fmt.Println("walk")
}

func (c Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct{}

func (b Bird) Eat() {
	fmt.Println("worms")
}

func (b Bird) Move() {
	fmt.Println("fly")
}

func (b Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct{}

func (s Snake) Eat() {
	fmt.Println("mice")
}

func (s Snake) Move() {
	fmt.Println("slither")
}

func (s Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	var request string
	animalia := make(map[string]Animal, 0)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Create an animal or query the created:")
	fmt.Println("Your commands must be space separated!!")
	fmt.Println("Enter X to exit.")
	for {
		fmt.Printf(">")
		scanner.Scan()
		request = scanner.Text()
		if request == "X" {
			break
		}
		data := strings.Split(strings.ToLower(request), " ")
		if len(data) != 3 {
			fmt.Println("You entered something unusual. Exiting ....")
		}
		action, name, param := data[0], data[1], data[2]
		if action == "newanimal" {
			switch param {
				case "cow":
					animalia[name] = Cow{}
					fmt.Println("Created It!")
				case "bird":
					animalia[name] = Bird{}
					fmt.Println("Created It!")
				case "snake":
					animalia[name] = Snake{}
					fmt.Println("Created It!")
				default:
					fmt.Println("That animal cannot be created.")
			}
		} else if action == "query" {
			switch param {
				case "eat":
					animalia[name].Eat()
				case "move":
					animalia[name].Move()
				case "speak":
					animalia[name].Speak()
				default:	
					fmt.Println("No such animal/behavior.")
			}
		}
	}	

}
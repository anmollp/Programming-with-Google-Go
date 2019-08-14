package main

import (
"fmt"
"bufio"
"os"
"strings"
)

type Animal struct {
	food string
	locomotion string
	noise string
}

func (a Animal) Eat() string{
	return a.food
}

func (a Animal) Move() string{
	return a.locomotion
}

func (a Animal) Speak() string{
	return a.noise
}

func Process(a *Animal, information string) string{
	if information == "eat" {
			return a.Eat()
		} else if information == "move" {
			return a.Move()
		} else if information == "speak" {
			return a.Speak()
		}
	return "No such information available"
}
func main() {
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}
	var request string
	fmt.Println("Enter animal and the information required with a space between them: ")
	fmt.Println("Otherwise enter X to exit")
	scanner := bufio.NewScanner(os.Stdin)
	for {
	fmt.Printf(">")
	scanner.Scan()
	request = scanner.Text()
	if request == "X"{
		break
	}
	data := strings.Split(request, " ")
	animal, information := data[0], data[1]
	animal = strings.ToLower(animal)
	information = strings.ToLower(information)
	switch animal {
		case "cow":
			fmt.Println(Process(&cow, information))
		case "bird":
			fmt.Println(Process(&bird, information))
		case "snake":
			fmt.Println(Process(&snake, information)) 
		default:
			fmt.Println("No such Animal")
		}
	}
}

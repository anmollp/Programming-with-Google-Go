package main

import "fmt"
import "sort"
import "strconv"

func main(){
	var capacity int = 3
	slice := make([]int, 0, capacity)
	for{
		var x string
		fmt.Println("Enter an integer:")
		_,err := fmt.Scan(&x)
		if x == "X" {
			return
		}
		if err != nil {
			fmt.Printf("Scanner error")
			return
		}
		number, err := strconv.Atoi(x)
		if err != nil {
			fmt.Println("You did not enter an Integer")
			return
		}
		slice = append(slice, number)
		sort.Ints(slice)
		fmt.Println(slice)
	}	

}

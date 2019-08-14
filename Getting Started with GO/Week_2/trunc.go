package main

import "fmt"

func main() {
	var x float32
	fmt.Printf("Enter a float number to truncate: ")
	fmt.Scan(&x)
	y := int(x)
	fmt.Printf("%d\n",y)
}

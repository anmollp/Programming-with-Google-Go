package main

import (
"fmt"
"os"
"bufio"
"strings"
"strconv"
)

func StringToSlice(s string) []int {
	slice := strings.Split(s, " ")
	intSlice:= make([]int, 0, 10)
	for _, element := range slice {
		num, err := strconv.Atoi(element)
		if err != nil {
			panic(err)
		}
		intSlice = append(intSlice, num)
	}
	return intSlice
}

func Swap(s []int, index int) {
	s[index] = s[index+1] + s[index]
	s[index+1] = s[index] - s[index+1]
	s[index] = s[index] - s[index+1]
}

func BubbleSort(s []int) {
	for i:=0; i<len(s); i++ {
		for j:=0; j<len(s)-i-1; j++ {
			if s[j] > s[j+1] {
				Swap(s, j)
			}
		}
	}
}

func main() {
	var intText string
	fmt.Println("Enter 10 integers, space sepearted!")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	intText = scanner.Text()
	
	intArray := make([]int, 0, 10)
	intArray = StringToSlice(intText)

	BubbleSort(intArray)

	fmt.Println("Sorted array: ", intArray)
}
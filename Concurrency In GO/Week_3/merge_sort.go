package main

import (
"fmt"
"bufio"
"os"
"strings"
"strconv"
)

func swap(subarray []int, index int) {
	subarray[index], subarray[index+1] = subarray[index+1], subarray[index]
}

func bubbleSort(subarray []int, c chan []int) {
	size := len(subarray)
	var swapped bool = false
	for i:=0; i<size; i++ {
		for j:=0;j<size-i-1;j++ {
		 	if subarray[j] > subarray[j+1] {
				swap(subarray, j)
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	c <- subarray
	fmt.Println("Sorted sub array:", subarray)
}

func merge(parent []int, child []int) []int {
	if len(parent) == 0 {
		return child
	}
	resultArray := make([]int, 0)
	i,j := 0,0
	for i < len(parent) && j < len(child) {
		if parent[i] < child[j] {
			resultArray = append(resultArray, parent[i])
			i++
		} else {
			resultArray = append(resultArray, child[j])
			j++
		}
	}

	for i < len(parent) {
		resultArray = append(resultArray, parent[i])
		i++
	}

	for j < len(child) {
		resultArray = append(resultArray, child[j])
		j++
	}
	return resultArray
}

func initiateSort(array []int, c chan []int) []int {
	size := len(array)
	windowSize := size/4
	start := 0
	stop := start + windowSize
	resultArray := make([]int, 0)
	for i:=1; i<4;i++ {
		go bubbleSort(array[start:stop], c)
		subarray := <- c
		resultArray = merge(resultArray, subarray)
		start = stop
		stop = start + windowSize
	}
	go bubbleSort(array[start:], c)
	subarray := <- c
	resultArray = merge(resultArray, subarray)
	return resultArray
}

func main() {
	fmt.Println("Enter space seperated integers:")
	var input string
	strArray := make([]string, 0)
	array := make([]int, 0)
	c := make(chan []int)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input = scanner.Text()
	strArray = strings.Split(input, " ")
	for i:=0; i < len(strArray); i++ {
		num, err := strconv.Atoi(strArray[i])
		if err != nil {
			panic("Not an integer")
		}
		array = append(array, num)
	}
	sortedArray := initiateSort(array, c)
	fmt.Println("Your Sorted array:",sortedArray)
}
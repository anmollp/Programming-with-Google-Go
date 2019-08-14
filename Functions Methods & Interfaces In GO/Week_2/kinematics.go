package main

import (
"fmt"
"math"
)

func GenDisplaceFn(a, v, d float64) func(float64) float64 {
	return func(t float64) float64 {
		s := (a * math.Pow(t,2)) / 2 + v * t+ d
		return s
	}
}

func main() {

	var acceleration, velocity, displacement, time float64
	fmt.Printf("Enter the acceleration: ")
	fmt.Scanf("%f", &acceleration)

	fmt.Printf("Enter the initial velocity: ")
	fmt.Scanf("%f", &velocity)

	fmt.Printf("Enter the initial displacement: ")
	fmt.Scanf("%f", &displacement)

	fmt.Printf("Enter the time after which displacement needs to be calculated: ")
	fmt.Scanf("%f", &time)

	dispOfTime := GenDisplaceFn(acceleration, velocity, displacement)
	fmt.Println("Final displacement:",dispOfTime(time))
}
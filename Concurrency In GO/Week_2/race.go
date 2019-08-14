package main

import (
 "fmt"
)

func routine1(x *int) {
	*x = *x + 2
}

func routine2(x *int) {
	*x = *x/ 2
}

func main() {
	var x int = 10
	go routine1(&x)
	go routine2(&x)
	fmt.Println(x)
}

// In this program there are two go routines namely routine1 and routine2
// These routines will be executed concurrently by the go runtime scheduler.
// To demonstrate the race condition happening , the integer variable x is used as 
// a communication between go routines.

// Since there are two goroutines which need to be scheduled concurrently,
// it is non deterministic and go run time scheduler can run any of the goroutines first.
// Hence there will be a non-deterministic interleaving of instructions present in these
// goroutines.  

// In the example above, if the goruntime scheduler schedules these routines say,
// firstly routine1 executes then routine2 and back to main. In such a case,
// The value of x will change from 10 to 12 in routine1 and from 12 to 6 in routine2.
// And finally in main it should print 6.
// But this is one among the many such interleavings of instructions of concurrent routines.
// If the order of execution was say, routine2->routine1->main, thrn the value of x 
// would be 7. Similarly for other combinations of routine schedules, we would get different 
// values of x. This is a race condition.

// Thus a race conditition is a situation where the outcome cannot be deterministic and
// depends on the non-deterministic interleaving of the instructions among different concurrent tasks/goroutines. 
// Also if there is no communication between concurrently running tasks there shall be no
// race conditions.

// To check what type of race condition might occur run the program with "-race" argument.
// Example: go run -race race.go
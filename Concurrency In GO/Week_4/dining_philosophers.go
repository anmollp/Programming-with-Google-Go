package main 

import (
	"fmt"
	"sync"
)

type ChopStick struct{
	sync.Mutex
}

type Philosopher struct {
	LeftCS *ChopStick
	RightCS *ChopStick
}

func AssignPermissions(c chan int){
	c <- 1
	c <- 1
}

func (p *Philosopher) Eat(id int, c chan int, wg *sync.WaitGroup) { 
	//Get Permission
	<- c
	p.LeftCS.Lock()
	p.RightCS.Lock()
	fmt.Println("Starting to eat ", id)
	fmt.Println("Finishing eating ", id)
	p.LeftCS.Unlock()
	p.RightCS.Unlock()
	wg.Done()
	c <- 1
}

func main() {
	permChannel := make(chan int, 2)
	chopsticks := make([]*ChopStick, 5)
	philosophers := make([]*Philosopher, 5)

	var wg sync.WaitGroup

	for i:=0;i<len(chopsticks);i++ {
		chopsticks[i] = new(ChopStick)
	}

	for i:=0;i<len(philosophers);i++ {
		philosophers[i] = &Philosopher{chopsticks[i], chopsticks[(i+1) % 5]}
	}
	
	for i:=0; i<3;i++ {
		wg.Add(len(philosophers))
		go AssignPermissions(permChannel)
		for p:=0; p<len(philosophers); p++ {
			go philosophers[p].Eat(p+1, permChannel, &wg)
		}
		wg.Wait()
	}
}


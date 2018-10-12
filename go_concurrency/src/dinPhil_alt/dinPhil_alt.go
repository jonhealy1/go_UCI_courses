package main

import "sync"
import "fmt"

//Chopstick represent chopstic as a mutex
type Chopstick struct {
	sync.Mutex
}

//Philosopher represents a dining philosopher
type Philosopher struct {
	leftCS, rightCS *Chopstick
	number          int
}

const eatNumber = 3

var wg sync.WaitGroup

func (p Philosopher) eat(doneEating, canEat chan int) {
	// Eat only 3 times
	for i := 0; i < eatNumber; i++ {
		// ask host for permission
		<-canEat
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Println("starting to eat ", p.number)
		fmt.Println("finishing eating ", p.number)
		p.rightCS.Unlock()
		p.leftCS.Unlock()
		// communicate to host finished
		doneEating <- 1
	}
	wg.Done()
}

func host(doneEating, canEat chan int) {
	// allow no more than two philosophers to eat at once
	canEat <- 1
	canEat <- 1
	for i := 0; i < 15; i++ {
		//Waiting for a philosopher to finish
		<-doneEating
		//Philospher finished eating, releasing forks
		canEat <- 1
	}
	wg.Done()

}

func main() {

	canEat := make(chan int, 2)
	doneEating := make(chan int, 2)
	//phil := make(chan int)

	//initialize Chopsticks and Philosophers
	chopsticks := make([]*Chopstick, 5)
	for i := 0; i < 5; i++ {
		chopsticks[i] = new(Chopstick)
	}
	philosophers := make([]*Philosopher, 5)
	wg.Add(1)
	go host(doneEating, canEat)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{chopsticks[i], chopsticks[(i+1)%5], i + 1}
		wg.Add(1)
		go philosophers[i].eat(doneEating, canEat)
	}
	wg.Wait()
}

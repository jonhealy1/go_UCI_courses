package main

/*
6 problems

sentences - a couple of sentences explaining how it could work in real computation
evaluation (tools/results)
compare/ contrast

					1	|	2
correctness
comprehesibility
performance
*/

import (
	"fmt"
	"sync"
)

// ChopS : chopsticks mutex
type ChopS struct{ sync.Mutex }

// Philo : philosopher needs left and right chopsticks to eat
type Philo struct {
	leftCS, rightCS *ChopS
}

func (p Philo) eat(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := 0; j < 3; j++ {
		if j%2 == 1 {
			p.leftCS.Lock()
			p.rightCS.Lock()
		} else {
			p.rightCS.Lock()
			p.leftCS.Lock()
		}
		fmt.Println("starting to eat: ", i)
		fmt.Println("finishing eating: ", i)
		p.rightCS.Unlock()
		p.leftCS.Unlock()
	}
}

func hosting(wg *sync.WaitGroup) {

	CSticks := make([]*ChopS, 5)

	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{CSticks[i], CSticks[(i+1)%5]}
	}
	var mutex = &sync.Mutex{}
	for i := 0; i < 3; i++ {
		if i < 2 {
			mutex.Lock()
			go philos[i].eat(i, wg)
			go philos[(i+2)%5].eat((i+2)%5, wg)
			mutex.Unlock()
		} else {
			mutex.Lock()
			go philos[(i+2)%5].eat((i+2)%5, wg)
			mutex.Unlock()
		}
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	go hosting(&wg)
	wg.Wait()
}

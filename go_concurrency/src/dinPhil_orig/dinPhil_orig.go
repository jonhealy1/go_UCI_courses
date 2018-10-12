package main

import (
	"fmt"
	"sync"
	"time"
)

// ChopS : chopsticks mutex
type ChopS struct{ sync.Mutex }

// Philo : philosopher needs left and right chopsticks to eat
type Philo struct {
	leftCS, rightCS *ChopS
}

func (p Philo) eat(i int) {
	for j := 0; j < 3; j++ {
		if j%2 == 1 {
			p.leftCS.Lock()
			p.rightCS.Lock()
		} else {
			p.rightCS.Lock()
			p.leftCS.Lock()
		}

		fmt.Println("starting to eat: ", i)

		p.rightCS.Unlock()

		p.leftCS.Unlock()
		//time.Sleep(1 * time.Second)
		fmt.Println("finishing eating: ", i)

	}
}

func hosting() {
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
			go philos[i].eat(i)
			go philos[(i+2)%5].eat((i + 2) % 5)
			mutex.Unlock()
		} else {
			go philos[(i+2)%5].eat((i + 2) % 5)
		}
	}
}

func main() {
	go hosting()
	time.Sleep(2 * time.Second)
}

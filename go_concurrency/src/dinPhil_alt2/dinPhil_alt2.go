package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type ChopS struct{ sync.Mutex }

type Philo struct {
	leftCS, rightCS *ChopS
	num, eats       int
	permission      chan int
	done            chan int
	wg              *sync.WaitGroup
}

func (p *Philo) eat() {
	for i := 0; i < 3; i++ {
		permission := <-p.permission
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Printf("starting to eat %d\n", p.num)
		time.Sleep(time.Duration(rand.Int()%5) * 100000000)
		p.eats++
		fmt.Printf("finishing eating %d\n", p.num)

		p.rightCS.Unlock()
		p.leftCS.Unlock()
		p.done <- permission
	}
	p.wg.Done()
}

func host(permissions chan int, done chan int) {
	permissions <- 1
	permissions <- 1

	for {
		permissions <- <-done
	}
}

func main() {
	var permissions = make(chan int, 2)
	var done = make(chan int, 2)

	wg := sync.WaitGroup{}
	wg.Add(5)

	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{CSticks[i], CSticks[(i+1)%5], i + 1, 0, permissions, done, &wg}
	}
	for i := 0; i < 5; i++ {
		go philos[i].eat()
	}
	go host(permissions, done)
	wg.Wait()
	for i := 0; i < 5; i++ {
		fmt.Println(philos[i].eats)
	}
}

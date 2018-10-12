package main

import (
	"fmt"
	"sync"
)

func crtclSection(mtx *sync.Mutex, wg *sync.WaitGroup, p string, i int, writer bool) {
	defer wg.Done()
	//r <- "reader"
	mtx.Lock()
	if writer == true {
		//mtx.Lock()
		for j := 0; j < 10; j++ {
			fmt.Println(p, i, j)
		}
		//mtx.Unlock()
	} else if writer == false {
		//mtx.Lock()
		for j := 0; j < 10; j++ {
			fmt.Println(p, i, j)
		}
		//mtx.Unlock()
		//fmt.Println(p)
	}
	mtx.Unlock()
	//return msg
}

func crtclSection2(wg *sync.WaitGroup, p string, i int, writer bool) {
	defer wg.Done()
	//r <- "reader"

	if writer == true {
		//mtx.Lock()
		for j := 0; j < 10; j++ {
			fmt.Println(p, i, j)
		}
		//mtx.Unlock()
	} else if writer == false {
		//mtx.Lock()
		for j := 0; j < 10; j++ {
			fmt.Println(p, i, j)
		}
		//mtx.Unlock()
	}

	//return msg
}

func main() {
	var mutex = &sync.Mutex{}
	var rdr sync.WaitGroup
	var wrt sync.WaitGroup
	rdr.Add(5)
	wrt.Add(5)
	//go hosting(&wg)
	var mutex2 = &sync.Mutex{}
	//r := make(chan string, 5)
	//w := make(chan bool, 5)
	//var msg string
	mutex2.Lock()
	for i := 0; i < 5; i++ {
		//go dance(m, men[i])

		go crtclSection(mutex, &rdr, "reading", i, false)
	}
	for i := 0; i < 5; i++ {
		//go dance(m, men[i])
		//mutex2.Lock()
		go crtclSection(mutex, &wrt, "writing", i, true)
		//mutex2.Unlock()
	}
	mutex2.Unlock()
	//msg := <-r
	//fmt.Println(msg)
	rdr.Wait()
	wrt.Wait()
}

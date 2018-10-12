package main

import (
	"fmt"
	"time"
)

var counter int

func dance(d chan string, person string) {
	//for i := 0; i < 5; i++ {

	d <- person
	time.Sleep(1 * time.Second)
	//msg := <-d
	//fmt.Println(msg)

}
func main() {
	men := []string{"bob", "al", "john", "phil", "sam"}
	women := []string{"lisa", "betty", "margaret", "phyllis", "donna"}
	m := make(chan string, 5)
	w := make(chan string, 5)
	//w := make(chan bool, 5)
	for i := 0; i < 5; i++ {
		//go dance(m, men[i])
		go dance(w, women[i])
	}
	for i := 0; i < 5; i++ {
		go dance(m, men[i])
		//go dance(w, women[i])
	}

	for i := 0; i < 5; i++ {
		msg := <-m
		fmt.Print(msg)
		msg2 := <-w
		fmt.Println(msg2)
	}
	/*
		msg := <-m
		fmt.Print(msg)
		msg = <-w
		fmt.Println(msg)
		msg = <-m
		fmt.Print(msg)
		msg = <-w
		fmt.Println(msg)
		msg = <-m
		fmt.Print(msg)

		msg = <-w
		fmt.Println(msg)
		msg = <-m
		fmt.Print(msg)
		msg = <-w
		fmt.Println(msg)
		msg = <-m
		fmt.Print(msg)
		msg = <-w
		fmt.Println(msg)
		//msg = <-d*/
	//time.Sleep(2 * time.Second)
}

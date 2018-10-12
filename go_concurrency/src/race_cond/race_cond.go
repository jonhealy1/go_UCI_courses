package main

import (
	"fmt"
	"time"
)
/*
A race condition happens when output is dependent on the 
sequence of uncontrollable events. The race condition in 
my code occurs because two different go routines are 
calling identical functions and updating the same 
variable in the process. If you run my code maybe 5 times 
you will see that the order that the routines run is not 
always the same. This can occur because even though go 
routine 1 should go first it doesn’t always because of 
scheduling and accessing shared memory.
*/
func main() {
	var a int
	a = 10
	
//	for j:=0; j<10; j++ {
	go addy(&a)
	go addy2(&a)
		
		
//	}
	time.Sleep(1 * time.Second)
}

func addy(a *int){
	
	*a = *a + 1
	fmt.Println("go 1",*a)

}
func addy2(a *int){
	
	*a = *a + 1
	fmt.Println("go 2",*a)

}
package main

import (
	"fmt"
	"time"
)

/**
	shared variable that will be update by the two go routines
 */
var counter int
/**
	function that will use to create go routines
 */
func f() {
	for i := 0; i < 10000; i++ {
		counter = counter +1
	}
}

func main() {
	/**
		creating go routines
	 */
	go f()
	go f()

	/**
		sleep main thread to wait go routines finish
	 */
	time.Sleep(time.Second)

	/**
	despite the counter in this moment should 20000
	It won't because we have a race caused by a write - write issue
	that means that both routines are update the var counter at the same time causing
	some operation be lost and very likely counter be less than 20.000
	*/
	fmt.Println(counter)
}

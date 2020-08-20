package main

import (
	"fmt"
	"sync"
	"time"
)

type ChopStick struct {
	sync.Mutex
}

type Philosopher struct {
	philosopherNumber int
	eatCount int
	leftChopStick, rightChopStick *ChopStick
}

func (p Philosopher) eat(c chan *Philosopher, wGroup *sync.WaitGroup){
	for i := 0; i < 3; i++ {
		c <- &p
		if p.eatCount < 3 {
			p.leftChopStick.Lock()
			p.rightChopStick.Lock()

			fmt.Println("Philosopher " , p.philosopherNumber , " is eating!")

			p.leftChopStick.Unlock()
			p.rightChopStick.Unlock()

			fmt.Println("Philosopher " , p.philosopherNumber , " is finished eating!")
			wGroup.Done()
		}
	}
}

func main(){
	var i int
	var waitGroup sync.WaitGroup
	c := make(chan *Philosopher,2)

	waitGroup.Add(15)

	ChopSticks := make([] *ChopStick,5)
	for i = 0; i < 5; i++ {
		ChopSticks[i] = new(ChopStick)
	}

	Philosophers := make([] *Philosopher,5)
	for i = 0; i < 5; i++ {
		Philosophers[i] = &Philosopher{i+1,0,ChopSticks[i],ChopSticks[(i+1)%5]}
	}

	go host(c)
	for i = 0; i < 5; i++ {
		go Philosophers[i].eat(c,&waitGroup)
	}
	waitGroup.Wait()
}

func host(c chan *Philosopher) {
	for {
		if len(c)==2 {
			<- c
			<- c
			time.Sleep(20 * time.Millisecond)
		}
	}
}

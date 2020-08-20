package main

import (
	//"bufio"
	"fmt"
    "sync"
	//"os"
	//"sort"
	//"strings"
	//"strconv"
	//"github.com/astaxie/beego/validation"
	//"log" //log stuff
	//"text/tabwriter"
)

type ChopS struct{ sync.Mutex }

type Philo struct {
	leftCS, rightCS *ChopS
}

func (p Philo) eat(num int) {
   for {
      p.leftCS.Lock()
      p.rightCS.Lock()

        fmt.Println("starting to eat", num+1)
        fmt.Println("finishing eating", num+1)

      p.rightCS.Unlock()
      p.leftCS.Unlock()
   }
}



func main() {
    
    CSticks := make([]*ChopS, 5)
    for i := 0; i < 5; i++ {
       CSticks[i] = new(ChopS)
    }
    
    philos := make([]*Philo, 5)
    for i := 0; i < 5; i++ {
       philos[i] = &Philo{CSticks[i], CSticks[(i+1)%5]}
    }
   
    for i := 0; i < 5; i++ {
         for j := 0; j < 5; j++ {
             go philos[i].eat(j)
        
         }
    }

    
}//end main
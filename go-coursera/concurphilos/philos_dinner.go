package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)

type ChopS struct {
	sync.Mutex
}

type Philo struct {
	idPhilo int
	leftCS, rightCS *ChopS
}

func (p Philo) eat(wg *sync.WaitGroup) {
	defer wg.Done()
	for i:=0; i<3; i++ {
		p.leftCS.Lock()
		p.rightCS.Lock()
		
		fmt.Printf("\nStarting to eat %d", p.idPhilo)
		time.Sleep(500 * time.Millisecond)
		
		p.leftCS.Unlock()
		p.rightCS.Unlock()
		
		fmt.Printf("\nFinishing eating %d", p.idPhilo)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	// Initialization for sticks & philos
	CSticks := make([]*ChopS, 5)
	
	for i:=0; i<5; i++ {
		CSticks[i] = new(ChopS)
	}
	
	philos := make([]*Philo, 5)
	
	for i:=0; i<5; i++ {
		philos[i] = &Philo{i+1, CSticks[i], CSticks[(i+1)%5]}
	}
	
	
	rand.Shuffle(len(philos), func(i, j int) {
		philos[i], philos[j] = philos[j], philos[i]
	})
	
	var wg sync.WaitGroup
	// Start the dinning
	for _, philo := range philos {
		wg.Add(1)
		go philo.eat(&wg)
	}
	wg.Wait()
}
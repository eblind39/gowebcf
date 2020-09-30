package main

import (
	"fmt"
	"strconv"
	"sync"
)

//ChopS struct
type ChopS struct{ sync.Mutex }

//Philo struct
type Philo struct {
	num             int
	leftCS, rightCS *ChopS
}

var wg sync.WaitGroup
var on sync.Once
var fullcounter int = 0
var philos []*Philo = make([]*Philo, 5)
var csticks []*ChopS = make([]*ChopS, 5)

func (p *Philo) eat() {
	p.leftCS.Lock()
	p.rightCS.Lock()
	fmt.Println("starting to eat <" + strconv.Itoa(p.num) + ">")
	p.leftCS.Unlock()
	p.rightCS.Unlock()
	fmt.Println("finishing eating <" + strconv.Itoa(p.num) + ">")
	wg.Done()
}

func setup() {
	wg.Add(10)
	for i := 0; i < 5; i++ {
		csticks[i] = new(ChopS)
		wg.Done()
	}
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{i + 1, csticks[i], csticks[(i+1)%5]}
		wg.Done()
	}
	wg.Wait()
	return
}

func main() {
	on.Do(setup)
	for i := 0; i < 3; i++ {
		wg.Add(5)
		for i := 0; i < 5; i++ {
			go philos[i].eat() //philosophers eating
		}
		wg.Wait()
	}
}

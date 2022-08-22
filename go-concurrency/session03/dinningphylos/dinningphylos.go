package main

import (
	"github.com/fatih/color"
	"sync"
	"time"
)

// The dinning philosophers problem is well known in computer science circles.
// Five philosophers, numbered from 0 to 4, live in a house where the
// table is laid for them; each philosopher has their own place at the table.
// Their only difficulty - besides those of philosophy - is that the dish
// served is a very difficult kind of spaghetti which has to be eaten with
// two forks. There are two forks next to each plate, so that presents no
// difficulty. AS a consequence, however, this means that no two neighbours
// may be eating simultaneously.

// constants
const hunger = 3

// variables - philosophers
var philosophers = []string{"Plato", "Socrates", "Aristotle", "Pascal", "Locke"}
var wg sync.WaitGroup
var sleepTime = 1 * time.Second
var eatTime = 2 * time.Second
var thinkTime = 1 * time.Second

func diningProblem(philosopher string, leftFork, rightFork *sync.Mutex) {
	defer wg.Done()

	// print a message
	color.Green("%s is seated", philosopher)
	time.Sleep(sleepTime)

	for i := hunger; i > 0; i-- {
		color.Yellow("%s is hungry.", philosopher)
		time.Sleep(sleepTime)

		// lock both forks
		leftFork.Lock()
		color.Green("\t%s picked up the fork to his left.", philosopher)
		rightFork.Lock()
		color.Green("\t%s picked up the fork to his right.", philosopher)

		// print a message
		color.Green("\t> %s has both forks, and is eating.", philosopher)
		time.Sleep(eatTime)

		// give the philosopher some time to think
		color.Cyan("\t%s is thinking", philosopher)
		time.Sleep(thinkTime)

		// unlock the mutexes
		leftFork.Unlock()
		color.Cyan("\t %s put down the fork on his left.", philosopher)
		rightFork.Unlock()
		color.Cyan("\t %s put down the fork on his right.", philosopher)
		time.Sleep(sleepTime)
	}

	// print out done message
	color.Blue("%s is satisfied.\n", philosopher)
	time.Sleep(sleepTime)

	color.Cyan("%s has left the table.\n", philosopher)
}

func main() {
	// print intro
	color.Cyan("The Dinning Philosophers Problem")
	color.Cyan("--------------------------------")

	wg.Add(len(philosophers))

	forkLeft := &sync.Mutex{}

	// spawn one go routine for each philosopher
	for i := 0; i < len(philosophers); i++ {
		// create a mutex for the right fork
		forRight := &sync.Mutex{}
		// spawn a goroutine
		go diningProblem(philosophers[i], forkLeft, forRight)

		forkLeft = forRight
	}

	wg.Wait()

	color.Cyan("The table is empty.")
}

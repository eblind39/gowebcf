package main

import (
	"fmt"
	"sync"
)

var msg string

func updateMessage(s string, wg *sync.WaitGroup, mux *sync.Mutex) {
	defer wg.Done()
	mux.Lock()
	msg = s
	fmt.Println(msg)
	mux.Unlock()
}

func main() {
	var wg sync.WaitGroup
	var mux sync.Mutex
	msg = "Hello, world!"

	wg.Add(2)
	go updateMessage("Hello, universe!", &wg, &mux)
	go updateMessage("Hello, cosmos!", &wg, &mux)
	wg.Wait()

	fmt.Println(msg)
}

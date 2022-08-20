package main

import (
	"strings"
	"sync"
	"testing"
)

func Test_updateMessage(t *testing.T) {
	var wg sync.WaitGroup
	var mux sync.Mutex
	const newMsg = "Goodbye, bad habits!"
	msg = "Hello, world!"

	wg.Add(2)
	go updateMessage("x", &wg, &mux)
	go updateMessage(newMsg, &wg, &mux)
	wg.Wait()

	if strings.Contains(msg, newMsg) || strings.Contains("x", newMsg) {
		t.Error("Incorrect value in msg")
	}
}

/*
 * Use the next method to build the program (using the race detector flag)
 *
 * go build -race .\race_cond.go
 *
 * and then execute race_cond.exe
 *
 * You'll get an output similar to:
 * 
 * ==================
 * WARNING: DATA RACE
 * Write at 0x0000007c1ba8 by goroutine 8:
 *  main.sumSlice()
 */
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	// resslice is a variable setted by all goroutines.
	resslice float64
	// wg is used to wait for the program to finish.
	wgIns sync.WaitGroup
)

func main() {
	var slc = make([]int, 5, 5)
	slc = []int{1, 2, 3, 4, 5}

	wgIns.Add(2)

	/*
	 * Each goroutine overwrites the work of the other. 
	 * Each goroutine makes its own copy of the  resslice variable and
	 * then is swapped out for the other goroutine.
	 */
	go sumElemsSlice(slc...)
	go sumElemsSlice(slc...)

	wgIns.Wait()

	fmt.Println("The result is: ", resslice)
}

func sumElemsSlice(vals ...int) {
	var sumElems int
	
	defer wgIns.Done()
	
	for _, val := range vals {
		sumElems += val

		runtime.Gosched()

		// Two goroutines are created trough this function, 
		// which reads and writes to the package variable resslice,
		// which is our shared resource in this example.
		// The next line represents the race condition between both goroutines
		resslice = float64(sumElems )
	}
}
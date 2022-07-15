package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Número de CPUs init: ", runtime.NumCPU())
	fmt.Println("Número de gorutinas init: ", runtime.NumGoroutine())
	var wg sync.WaitGroup
	res01 := make(chan float64)
	res02 := make(chan float64)

	wg.Add(2)
	go func() {
		res01 <- gorut01(360)
		wg.Done()
	}()
	go func() {
		res02 <- gorut02(720)
		wg.Done()
	}()

	fmt.Println("Número de CPUs after run gorut: ", runtime.NumCPU())
	fmt.Println("Número de gorutinas run gorut: ", runtime.NumGoroutine())

	fmt.Println("res01", <-res01)
	fmt.Println("res02", <-res02)

	wg.Wait()
	close(res01)
	close(res02)
}

func gorut01(number float64) float64 {
	for i := 0; i < 10000; i++ {
		number = math.Sin(number)
		number = math.Asin(number)
	}
	return number
}

func gorut02(number float64) float64 {
	for i := 0; i < 10000; i++ {
		number = math.Cos(number)
		number = math.Acos(number)
	}
	return number
}

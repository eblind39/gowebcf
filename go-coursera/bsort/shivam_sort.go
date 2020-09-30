package main

import "fmt"

func Swap(slice []int, i int){
	slice[i], slice[i+1] = slice[i+1], slice[i]
}

func BubbleSort(slice []int){
	for i:=0; i<10; i++{

		for j:=0; j<10-i-1; j++{
			if slice[j]>slice[j+1]{
				// Swap func call
				Swap(slice, j)
			}
		}
	}
	fmt.Print(slice)
}

func main(){
	// init
	var slice []int
	var elm int

	// not more than 10 inputs are allowed
	for i:=0; i<10; i++{
		fmt.Scan(&elm)
		slice = append(slice, elm)
	}
	// Sort func call
	BubbleSort(slice)
}
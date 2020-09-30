package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var unSortedNumbers []int
	unSortedNumbers, err := getDataFromUsers()
	if err != nil{
		fmt.Println("Issue with getting input numbers from users")
	}
	BubbleSort(unSortedNumbers)
	fmt.Println(unSortedNumbers)
}

func BubbleSort(unSortedNumbers []int)  {
	n := len(unSortedNumbers)

	for i:=0;i<n;i++{
		for j:=0;j<n-i-1;j++{
			if unSortedNumbers[j] > unSortedNumbers[j+1]{
				Swap(unSortedNumbers, j)
			}
		}
	}
}
func Swap(unSortedNumbers []int, index int)  {
	temp:= unSortedNumbers[index+1]
	unSortedNumbers[index+1]= unSortedNumbers[index]
	unSortedNumbers[index]=temp
}
func getDataFromUsers()([]int, error)  {
	fmt.Println("Please enter numbers separated by new line (press enter after every number you enter), press x to stop entering numbers")
	var input string
	var unSortedNumbers []int
	for {
		if _,err:=fmt.Scan(&input);err != nil{
			return nil, err
		}
		if strings.ToLower(input) == "x"{
			return unSortedNumbers, nil
		}
		if v, err := strconv.Atoi(input); err != nil{
			return nil, err
		}else{
			unSortedNumbers = append(unSortedNumbers, v)
		}

	}
}

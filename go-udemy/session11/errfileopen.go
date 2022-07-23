package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fhndrl, err := os.Open("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer fhndrl.Close()

	bs, err := ioutil.ReadAll(fhndrl)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bs))
}

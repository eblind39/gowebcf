package main

import "fmt"

type mapstr map[string]string

func main() {
	ms := make(map[string]string)
	mstr1 := mapstr{}
	mstr2 := make(mapstr)

	ms["one"] = "one"
	mstr1["two"] = "two"
	mstr2["three"] = "three"

	fmt.Println(ms["one"], mstr1["two"], mstr2["three"])
}

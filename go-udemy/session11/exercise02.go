package main

import "fmt"

type CustomErr struct {
	info string
}

func (ce CustomErr) Error() string {
	return fmt.Sprintf("An error has ocurred: %v", ce.info)
}

func main() {
	e1 := CustomErr{
		info: "I need to sleep more",
	}

	fooerr(e1)
}

func fooerr(e error) {
	fmt.Println("foo has ran - ", e, "\n", e.(CustomErr).info)
}

package main

import (
	"fmt"
	"log"
)

type raizError struct {
	lat  string
	long string
	err  error
}

func (re raizError) Error() string {
	return fmt.Sprintf("error matemático: %v %v %v", re.lat, re.long, re.err)
}

func main() {
	_, err := raiz(-10.23)
	if err != nil {
		log.Println(err)
		return
	}
}

func raiz(f float64) (float64, error) {
	if f < 0 {
		// Se podría usar:
		// e := errors.New("Necesito dormir más")
		// O bien:
		e := fmt.Errorf("Necesito dormir más - el valor era: %v", f)
		return 0, raizError{"50.2289 N", "99.4656 W", e}
	}
	return 42, nil
}

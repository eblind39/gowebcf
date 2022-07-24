package architecture

import (
	"fmt"
	"testing"
)

// mocking storage
type Db map[int]Persona

func (m Db) Save(n int, p Persona) {
	m[n] = p
}
func (m Db) Retrieve(n int) (p Persona) {
	return m[n]
}

func TestPut(t *testing.T) {
	mdb := Db{}
	p := Persona{
		First: "Steven",
	}

	psm := NewPersonService(mdb)

	// psm.Put(2, p) // Fail
	psm.Put(1, p) // Pass

	got := mdb.Retrieve(1)

	if got != p {
		t.Fatalf("Expected %v, got %v", p, got)
	}
}

func ExamplePut() {
	mdb := Db{}
	p := Persona{
		First: "Steven",
	}
	psm := NewPersonService(mdb)
	psm.Put(1, p) // Pass
	got := mdb.Retrieve(1)
	fmt.Println(got)
	//Output: {Steven}
}

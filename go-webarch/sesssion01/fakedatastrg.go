package main

import "fmt"

type personn struct {
	first string
}

type mongo map[int]personn
type postg map[int]personn

func (m mongo) save(n int, p personn) {
	m[n] = p
}
func (m mongo) retrieve(n int) (p personn) {
	return m[n]
}

func (pg postg) save(n int, p personn) {
	pg[n] = p
}
func (pg postg) retrieve(n int) (p personn) {
	return pg[n]
}

type accesor interface {
	save(n int, p personn)
	retrieve(n int) personn
}

type personService struct {
	a accesor
}

func (ps personService) put(n int, p personn) {
	ps.a.save(n, p)
}

func (ps personService) get(n int) (personn, error) {
	p := ps.a.retrieve(n)
	if p.first == "" {
		return personn{}, fmt.Errorf("No person with n of %d", n)
	}
	return p, nil
}

func main() {
	dbm := mongo{}
	dbp := postg{}

	p1 := personn{
		first: "Juri",
	}

	p2 := personn{
		first: "Cooper",
	}

	// Using a service
	psm := personService{
		a: dbm,
	}
	psm.put(1, p1)
	psm.put(2, p2)
	fmt.Println(psm.get(1))
	fmt.Println(psm.get(3))

	// or store in another data store
	psp := personService{
		a: dbp,
	}
	psp.put(1, p1)
	psp.put(2, p2)
	fmt.Println(psp.get(1))
	fmt.Println(psp.get(2))
}

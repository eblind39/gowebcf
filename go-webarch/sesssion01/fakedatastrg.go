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

func put(a accesor, n int, p personn) {
	a.save(n, p)
}
func get(a accesor, n int) personn {
	return a.retrieve(n)
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

	put(dbm, 1, p1)
	put(dbm, 2, p2)
	fmt.Println(get(dbm, 1))
	fmt.Println(get(dbm, 2))

	// or store in another data store
	put(dbp, 1, p1)
	put(dbp, 2, p2)
	fmt.Println(get(dbp, 1))
	fmt.Println(get(dbp, 2))
}

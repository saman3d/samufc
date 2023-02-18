package common

type Location struct {
	Street    Street
	Apartment Apartment
	City      City
	State     State
	Country   Country
}

type Apartment struct {
	Number uint64
	Name   string
}

type Street struct {
	Number uint64
	Name   string
}

type City struct {
	ID   uint64
	Name string
}

type State struct {
	ID   uint64
	Name string
}

type Country struct {
	ID      uint64
	Name    string
	TriCode string
}

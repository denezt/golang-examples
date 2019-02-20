package main

import (  
    "fmt"
)

type hamster struct {  
    name string
    breed string
    age int
    size int
}

func New(name string, breed string, age int, size int) hamster {
	h := hamster {name, breed, age, size}
	return h
}

func (h hamster) getMeasurements() {  
    fmt.Printf("%s is a %s has reach %d of maxsize.", h.name, h.breed, (h.size))
}

func main() {
	//h := hamster("Bruce Lee", "Chinese Dwarf Hamster", 2, 5)
	h := new hamster()
	h.getMeasurements()
}


package main

import (
	"fmt";
	"./test"
);

func main() {
	// Create an object of type Dog from the animals package.
	dog := animals.Hamster {
		Animal: animals.Animal{
			Name: "Bruce Lee",
			Age:  1,
		},
		CurrentSize: 10,
	}

	fmt.Printf("Output: %#v\n", dog)
}


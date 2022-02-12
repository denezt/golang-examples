package main
import "fmt"

func main() {
	// A Map is a collection of key value pairs
	// Created with varName := make(map[keyType] valueType)
	presAge := make(map[string] int)
	presAge["Narendra Modi"] = 42
	fmt.Println(presAge["Narendra Modi"])
	// Get the number of items in the Map
	fmt.Println(len(presAge))
	// The size changes when a new item is added
	presAge["Rahul Gandhi"] = 43
	fmt.Println(len(presAge))
	// We can delete by passing the key to delete
	delete(presAge, "Rahul Gandhi")
	fmt.Println(len(presAge))
}

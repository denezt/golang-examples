package animals

// Animal represents information about all animals.
type Animal struct {
	Name string
	Age  int
}

// Hamster represents information about hamsters.
type Hamster struct {
	Animal
	CurrentSize int
}

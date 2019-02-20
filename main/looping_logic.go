package main
import "fmt"

func main() {
	// For loops
	i := 1
	for i <= 10 {
		fmt.Println(i)
		// Shorthand for i = i + 1
		i++
	}
	// Relational Operators include ==, !=, <, >, <=, and >=
	// You can also define a for loop like this, but you need semicolons
	for j := 0; j < 5; j++ {
		fmt.Println(j);
	}
}

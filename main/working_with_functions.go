package main
import "fmt"

func main () {
	fmt.Println("5 + 4 = ", add(5,4))
	fmt.Println(subtract(1,2,3,4,5))
}

func add(a,b int) int {
	return a+b
}

func subtract(args... int) int {
	sub := 0
	for _, value := range args {
		sub -= value
	}
return sub
}

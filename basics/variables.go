package main
import "fmt"

func main () {
	var x int = 5 //declaration of variable x
	var (
	a = 10
	b = 15 //declaration of multiple variables
	)
	y := 7 // shorthand declaration of variables
	const pi float64 = 3.14272345 //declaration of constants
	var name string = "Will Smiths" //declaration of a string
	fmt.Println("Address: ",&a,"Value: ",a)
	fmt.Println(&b)
	fmt.Println(&x)
	fmt.Println(&y)
	fmt.Println(pi)
	fmt.Println(&name)
}
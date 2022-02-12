package main
import "fmt"

func main () {
	var x int = 5 //declaration of variable x
	var var1, var2, var3 = "foo", 100, "bar"
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
	fmt.Printf("var1 is type: %T value: ", var1)
	fmt.Println(var1)
	fmt.Printf("var2 is type: %T value: ", var2)
	fmt.Println(var2)
	fmt.Printf("var3 is type: %T value: ", var3)
	fmt.Println(var3)
}

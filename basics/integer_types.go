package main

import "fmt"

func main(){
	fmt.Println("Running, integers example...")
	// Calculate during runtime
	fmt.Println(1 + 1)
	// No calculation, this is of data type string
	fmt.Println("1 + 1")

	// Other Integer Data types
	// uint8
	// unsigned 8-bit integer (0 - 255)

	// uint16
	// unsigned 16-bit integer (0 - 65535)

	// uint32
	// unsigned 32-bit integer (0 - 4294967295)

	// int8
	// signed 8-bit integer (-127 - 127)

	// int16
	// signed 16-bit integer (-32768 - 32768)

	// int32
	// signed 32-bit integer (-2147483648 - 2147483648)

	// int64
	// signed 32-bit integer (-9223372036854775808 - 9223372036854775808)

	// Floating Point
	// floatN
	var (
		// float32
		a float32 = 4.0000 / 5.00045
		// float64
		b float64 = 4.0000 / 5.00045
	)
	fmt.Printf("%T : %f\n", a, a)
	fmt.Printf("%T : %f\n", b, b)
	// complexN
	var (
		// complex64
		x complex64 = 4.0000 + 5.00045
		// complex128
		y complex128 = 4.0000 + 5.00045
	)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	// Checking data type when it is not specified
	var myvar = 0.0
	fmt.Printf("%T\n", myvar)
}

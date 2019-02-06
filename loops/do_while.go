package main

import "fmt"

func main() {
    valid := true
    i := 0

    // This loop continues while "valid" is true.
    for valid {

        // If i equals 3, set "valid" to false.
        if i == 3000 {
            valid = false
        }
        fmt.Println(i)
        i++
    }
}


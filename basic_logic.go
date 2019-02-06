
package main
import "fmt"
func main() {
// If Statement
age := 15
if age >= 16 {
fmt.Println("Adult")
} else {
fmt.Println("Not an adult")
}
// You can use else if perform different actions, but once a match
// is reached the rest of the conditions aren't checked
if age >= 16 {
fmt.Println("in school")
} else if age >= 18 {
fmt.Println("in college")
} else {
fmt.Println("probably dead")
}
// Switch statements are used when you have limited options
switch age {
case 16: fmt.Println("Go Drive")
case 18: fmt.Println("Go Vote")
default: fmt.Println("Go Have Fun")
}
}

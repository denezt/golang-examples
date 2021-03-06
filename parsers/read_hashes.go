package parser

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    // Open the file.
    f, _ := os.Open("render_hash.hctxx")
    // Create a new Scanner for the file.
    scanner := bufio.NewScanner(f)
    // Loop over all lines in the file and print them.
    for scanner.Scan() {
        line := scanner.Text()
        fmt.Println(line)
    }
}

package parser

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func main() {
	in := `{"firstName":"George","lastName":"Washington"}`
	bytes := []byte(in)

	var p Person
	err := json.Unmarshal(bytes, &p)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 3; i++ {
		fmt.Printf("%+v\n", p)
	}
}

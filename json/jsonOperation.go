package main

import (
	"encoding/json"
	"fmt"

	"github.com/sanity-io/litter"
)

type sval struct {
	Name string
}
type (
	FullJson struct {
		Address Address
		Name    string
		Age     int
		Pets    []Pet
	}
	Address struct {
		Line1  string
		Line2  string
		Postal string
	}
	Pet struct {
		Name  string
		Kind  string
		Age   int
		Color string
	}
)

func main() {
	input := `
		{
		"name":"jai pal"
		}
`
	complexJson := `{
  "name": "James Peterson",
  "age": 37,
  "address": {
    "line1": "Block 78 Woodgrove Avenue 5",
    "line2": "Unit #05-111",
    "postal": "654378"
  },
  "pets": [
    {
      "name": "Lex",
      "kind": "Dog",
      "age": 4,
      "color": "Gray"
    },
    {
      "name": "Faye",
      "kind": "Cat",
      "age": 6,
      "color": "Orange"
    }
  ]
}
`

	var target sval
	err := json.Unmarshal([]byte(input), &target)
	if err != nil {
		fmt.Println("It's have json error")
	}

	var dog FullJson
	dogErr := json.Unmarshal([]byte(complexJson), &dog)

	if dogErr != nil {
		fmt.Println("It's have json error in dog")
	}
	// for k, v := range target {
	// 	fmt.Println(k, ":", v)
	// }

	fmt.Println(target.Name)

	litter.Dump(dog)

}

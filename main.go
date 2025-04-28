package main

import (
	"fmt"

	"github.com/mrangel-jr/ontraport/utils"
)

func main() {

	input := []byte(`{
	"one": {
			"two": 3,
			"four": [5,6,7]
	},
	"eight": {
			"nine": {
					"ten": 11
					}
					}
					}`)

	result, err := utils.UnmarshalJSON(input)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("result:", result)
	reversed, err := utils.MarshalJSON(result)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	fmt.Printf("%s: %s\n", "reversed", string(reversed))
}

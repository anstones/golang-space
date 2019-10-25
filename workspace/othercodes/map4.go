package main

import "fmt"

func main() {
	dict := make(map[string]string)

	dict = map[string]string{"red": "xxxxx", "orange": "@@@"}

	fmt.Printf("dict: %+v\n", dict)

	value, exists := dict["red"]
	if exists {
		fmt.Println(value)
	}

	delete(dict, "red")

	value, ok := dict["red"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("key has delete")
	}

}

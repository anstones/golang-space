package main

import (
	"fmt"
)

func Yes() string {
	fmt.Println("yes")
	return "yes"
}

func main() {
	var x string
	x, _ := Yes()
	// x, _ = Yes()
	// x, y := Yes()
}

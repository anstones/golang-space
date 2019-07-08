package main

import (
	"fmt"
	"unicode/utf8"
)

func printBytes(s string) {
	for i:= 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}


func printChars(s string) {
	runes := []rune(s)
	for i:= 0; i < len(runes); i++ {
		fmt.Printf("%c ",runes[i])
	}
}

func length(s string) {
	fmt.Printf("length of %s is %d\n", s, utf8.RuneCountInString(s))
}

func main() {
	name := "Hello World"
	//name := "Señor"
	printBytes(name)
	fmt.Printf("\n")
	printChars(name)
	fmt.Printf("\n")
	length(name)
}
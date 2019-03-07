package main

import (
	"fmt"
)

func main() {
	setmap := map[string]string{"china":"beijing"}
	setmap = map[string]string{"city":"35"}


    for contury := range setmap{
		fmt.Println(contury,"首都是",setmap[contury])
	}

	city, ok := setmap["china"]
	fmt.Println(city)
	fmt.Println(ok)
}
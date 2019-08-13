package main

import (
	"fmt"
	"strconv"
	"time"
)

const LOOP = 10000

var num int64 = 10

func main() {

	starttime := time.Now()

	for i := 0; i < LOOP; i++ {
		fmt.Sprintf("%d", i)
	}

	fmt.Printf("fmt.Sprintf taken: %v\n", time.Since(starttime))

	starttime = time.Now()
	for i := 0; i < LOOP; i++ {
		strconv.Itoa(i)
	}
	fmt.Printf("strconv.FormatInt taken: %v\n", time.Since(starttime))

}

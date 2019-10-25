package main

import (
	"fmt"
)

func main() {
	for m:=1;m<10;m++{
		for n:=1;n<=m;n++{
			fmt.Printf("%dX%d=%2d ",n,m,m*n)
		}
		fmt.Println("")
	}
}

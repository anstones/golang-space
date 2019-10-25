package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func checkerror(err error){
	if err != nil{
		log.Fatal("File reading error", err)
		return
	}
}

func main()  {
	data, err := os.Open("C:\\Users\\Administrator\\Desktop\\oeasy\\data\\mine\\Project\\Go_Space\\src\\filehanding\\test1.txt")
	checkerror(err)

	defer func() {
		if err = data.Close(); err != nil{
			log.Fatal(err)
		}
	}()

	r := bufio.NewScanner(data)
	for r.Scan(){
		fmt.Println(r.Text())
	}

	err = r.Err()
	checkerror(err)
}

package chapter2

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main()  {
	message := "Away from keyboard https://goland.ord/"

	encodeMessage := base64.StdEncoding.EncodeToString([]byte(message))

	fmt.Println(encodeMessage)

	data,err := base64.StdEncoding.DecodeString(encodeMessage)

	if err != nil{
		log.Fatal(err)
	}else{
		fmt.Println(string(data))
	}
}

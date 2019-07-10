package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://www.baidu.com", strings.NewReader("key=value"))

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	req.Header.Add("User-Agent", "myClient")

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	fmt.Println(string(data))

}

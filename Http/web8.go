// 模拟客户端上传文件
package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func postFile(filename string, targeturl string) error  {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	//关键的一步操作
	fileWriter, err := bodyWriter.CreateFormField(filename)
	if err != nil{
		fmt.Println("error writing to buffer")
		return err
	}
	// 打开文件句柄操作
	fh, err := os.Open(filename)
	if err != nil{
		fmt.Println("error opening file")
		return err
	}
	//iocopy
	n, err := io.Copy(fileWriter, fh)
	if err != nil{
		return err
	}
	fmt.Printf("\n write %d err %v \n", n, err)

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(targeturl, contentType, bodyBuf)
	if err != nil{
		return err
	}
	defer  resp.Body.Close()

	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return err
	}

	fmt.Println(resp_body)
	fmt.Println(string(resp_body))
	return nil

}

func main()  {
	target_url := "http://localhost:9000/upload"
	filename := "C:\\Users\\Administrator\\Desktop\\haha.jpg"
	postFile(filename, target_url)
}

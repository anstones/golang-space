package main

import "fmt"

type HttpResponse struct {
	status_codes int
}

func (r *HttpResponse) validResopnse() {
	r.status_codes = 200
}

func (r HttpResponse) updateStatus() string {
	return fmt.Sprint(r)
}

func main() {
	var r1 HttpResponse //r1是值
	r1.validResopnse()
	fmt.Println(r1.updateStatus())

	r2 := new(HttpResponse) // r2是指针
	r2.validResopnse()
	fmt.Println(r2.updateStatus())
}

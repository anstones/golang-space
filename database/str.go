package main

// 字符串转换

import (
	"fmt"
	"strconv"
)
func main() {
	str := make([]byte, 0, 100)
	str = strconv.AppendInt(str, 456710, 10)
	str = strconv.AppendBool(str, false)
	str = strconv.AppendQuote(str, "abcdefg")
	str = strconv.AppendQuoteRune(str, '单')
	fmt.Println(string(str))

	a := strconv.FormatBool(false)
	b := strconv.FormatFloat(123.23, 'g', 12, 64)
	c := strconv.FormatInt(1234, 10)
	d := strconv.FormatUint(12345, 10)
	e := strconv.Itoa(1023)
	fmt.Println(a, b, c, d, e)

	f, err := strconv.ParseBool("false")
	if err != nil {
		fmt.Println(err)
	}
	g, err := strconv.ParseFloat("123.23", 64)
	if err != nil {
		fmt.Println(err)
	}
	h, err := strconv.ParseInt("1234", 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	i, err := strconv.ParseUint("12345", 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	j, err := strconv.Itoa("1023")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(f, g, h, i, j)
}
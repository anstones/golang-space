package main

import (
	"fmt"
)

type student struct {
	fristName string
	lastName string
	scoure string
	country   string
}

func filter(s []student,f func(student) bool)  []student{
	var r  []student
	for _,v := range s{
		if f(v) == true{
			r = append(r,v)
		}
	}
	return r
}
//
//func aa(s student) bool{
//	if s.scoure == "B" {
//		return true
//	}
//	return false
//}

func main() {
	s1 := student{
		fristName: "Naveen",
		lastName:  "Ramanathan",
		scoure:     "A",
		country:   "India",
	}
	s2 := student{
		fristName: "Samuel",
		lastName:  "Johnson",
		scoure:     "B",
		country:   "USA",
	}
	s := []student{s1, s2}
	f := filter(s, func(s student) bool {
		if s.scoure == "B" {
			return true
		}
		return false
	})
	fmt.Println(f)
}


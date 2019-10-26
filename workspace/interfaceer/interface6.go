package main

import (
	"fmt"
)

type notifier interface {
	notify()
}

type User struct {
	name  string
	email string
}

func (user *User) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", user.name, user.email)
}

type Admin struct {
	name  string
	email string
}

func (admin *Admin) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", admin.name, admin.email)
}

func sendNotification(n notifier) {
	n.notify()
}

func main() {
	bill := User{"bill", "bill@email.com"}
	lisa := Admin{"lisa", "lisa@email.com"}
	sendNotification(&bill)
	sendNotification(&lisa)
}

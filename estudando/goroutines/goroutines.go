package main

import "time"

func foo() {
	//fmt.Println("foo")
	println("foo")
}

func main() {
	go foo()

	time.Sleep(2 * time.Second)
	println("fim")
}

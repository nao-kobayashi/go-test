package main

import "fmt"

func calmDown() {
	p := recover()
	err, ok := p.(error)
	if ok {
		fmt.Println(err.Error())
	}
}

func hitTest() {
	defer calmDown()
	err := fmt.Errorf("there's an error")
	panic(err)
}

func main() {
	hitTest()
	fmt.Println("main end")
}

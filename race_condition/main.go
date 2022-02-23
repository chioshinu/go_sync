package main

import "fmt"

func A() {
	fmt.Print("A->")
}

func B() {
	fmt.Print("B")
}

func main() {
	i := 0
	for i < 100 {
		go A()
		go B()
		i++
	}
}

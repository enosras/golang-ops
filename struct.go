package main

import "fmt"

type Student struct {
	Name   string
	Age    int
	Wealth float64
	Talent string
}

func main() {
	fmt.Println("Hello")

	a := &Student{}
	a.Name = "enos"
	fmt.Println(a.Name)
}

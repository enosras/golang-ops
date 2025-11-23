package main

import "fmt"

//The purpose of this module is to experiment with function and input/output use cases

func main(){
	var a, b, c int

	fmt.Print("Enter First :")
	fmt.Scanf("%d", &a)
	fmt.Print("Enter Second :")
	fmt.Scanf("%d", &b)
	fmt.Println(a)
	c = a + b
	fmt.Println(c)
	fmt.Println("enos")
}

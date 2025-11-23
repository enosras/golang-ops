package main

import "fmt"

func main() {

	//fmt.Println("Enter the size of array please : ")
	//var size_arr int = 3
	//fmt.Scanf("%d", &size_arr)

	var arr [3]string = [3]string{"Enoh", "Amor", "Otis"}
	fmt.Println(arr[2])
	fmt.Println("Wait for it: ")
	//var length int = len(arr[])
	//fmt.Println(len(arr))


	
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])

	}

}

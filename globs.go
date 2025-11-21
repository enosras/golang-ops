package main

import (
	"fmt"
	"path/filepath"
	_ "strings"
)

func go_glob() {
	matches, err := filepath.Glob("*.go")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for _, match := range matches {
		fmt.Println(match)
	}
}

func py_glob() {
	matches, err := filepath.Glob("*.py")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for _, match := range matches {
		fmt.Println(match)
	}
}

func py_txt() {
	matches, err := filepath.Glob("*.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for _, match := range matches {
		fmt.Println(match)
	}
}

func main() {
	var i int
	fmt.Println("1. Go files")
	fmt.Println("2. Python files")
	fmt.Println("3. Text files")
	fmt.Print("Enter your choice: ")
	fmt.Scanln(&i)

	switch i {
	case 1:
		go_glob()
	case 2:
		py_glob()
	case 3:
		py_txt()
	default:
		fmt.Println("Invalid option")
	}

}

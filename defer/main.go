package main

import "fmt"

func main() {
	data := "jefri"

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("betulkan panic")
		}
	}()

	fmt.Println(data)
	panic("ada kesalahan")
	fmt.Println(data)
}

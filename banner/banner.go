package main

import "fmt"

func main() {
	banner("Hello, Gophers! 😎", 24)
	banner("Hello, Gophers!", 24)

	fmt.Println("len: ", len("Hello, Gophers! 😎"))
	fmt.Println("len: ", len("Hello, Gophers!"))
}

func banner(text string, width int) {
	padding := (width - len(text)) / 2
	for i := 0; i < padding; i++ {
		fmt.Print(" ")
	}
	fmt.Println(text)
	for i := 0; i < width; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}

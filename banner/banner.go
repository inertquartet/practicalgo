package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	banner("Hello, Gophers! 😎", 24)
	banner("Hello, Gophers!", 24)

	s := "Go"
	fmt.Println("len: ", len(s))
	for i, r := range s {
		fmt.Println(i, r)
		if i == 0 {
			fmt.Printf("%c of type %T\n", r, r) // rune (int32)
		}
	}

	b := s[0]
	fmt.Printf("%c of type %T\n", b, b) // byte (uint8)

	fmt.Println(isPalindrome("g"))
	fmt.Println(isPalindrome("go"))
	fmt.Println(isPalindrome("gog"))
	fmt.Println(isPalindrome("gogo"))
}

// isPalindrome("g") => true
// isPalindrome("go") => false
// isPalindrome("gog") => true
// isPalindrome("gogo") => false
func isPalindrome(s string) bool {
	// TODO: implementation goes here
	length := len(s)
	for i := 0; i < (length / 2); i++ {
		if s[i] != s[length-i-1] {
			return false
		}

	}
	return true
}

func banner(text string, width int) {
	//	padding := (width - len(text)) / 2 // BUG: len is in bytes, not runes
	padding := (width - utf8.RuneCountInString(text)) / 2
	for i := 0; i < padding; i++ {
		fmt.Print(" ")
	}
	fmt.Println(text)
	for i := 0; i < width; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}

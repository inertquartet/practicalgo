package main

import "fmt"

func main() {
	var i any
	// go < 1.18
	// v i interface{}

	i = 7
	fmt.Println(i)

	i = "hi"
	fmt.Println(i)

	s := i.(string) // type assertion
	fmt.Println("s: ", s)

	// comma ok idiom
	n, ok := i.(int) // will panic
	if ok {
		fmt.Println("n: ", n)
	} else {
		fmt.Println("not an int")
	}

	switch i.(type) { // type switch
	case int:
		fmt.Println("an int")
	case string:
		fmt.Println("a string")
	default:
		fmt.Printf("unknown type: %T\n", i)
	}
}

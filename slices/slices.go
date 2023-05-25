package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	/* 	var s []int                // s is a slice of ints
	   	fmt.Println("len", len(s)) // len is "nil safe"
	   	if s == nil {              // you can compare only a slice to nil
	   		fmt.Println("nil slice")
	   	}

	   	s2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	   	fmt.Printf("s2 = %#v\n", s2)

	   	s3 := s2[1:4] // slicing operation, half-open range
	   	fmt.Printf("s3 = %#v\n", s3)

	   	// fmt.Println(s2[:100]) // panic: runtime error: slice bounds out of range
	   	s3 = append(s3, 100)
	   	fmt.Printf("s3 (appended) = %#v\n", s3)
	   	fmt.Printf("s2 (appended) = %#v\n", s2) // s2 has been modified as well!!!
	   	fmt.Printf("s2: len=%d, cap=%d\n", len(s2), cap(s2))
	   	fmt.Printf("s3: len=%d, cap=%d\n", len(s3), cap(s3))

	   	// var s4 []int
	   	s4 := make([]int, 0, 1_000)
	   	for i := 0; i < 1_000; i++ {
	   		s4 = appendInt(s4, i)
	   	}
	   	fmt.Println("s4", len(s4), cap(s4)) */
	// s4[1001] = 7 // panic: runtime error: index out of range [1001] with length 1001

	fmt.Println(concat([]string{"A", "B"}, []string{"C", "D", "E"})) // [A B C D E]

	vs := []float64{2, 1, 3}
	fmt.Println(median(vs)) // 2
	vs = []float64{2, 1, 3, 4}
	fmt.Println(median(vs)) // 2.5
	fmt.Println(vs)         // [2 1 3 4]

	fmt.Println(median(nil))
	fmt.Println(reflect.TypeOf(2))
}

func median(values []float64) (float64, error) {
	if len(values) == 0 {
		return 0, fmt.Errorf("median of empty slice is not defined")
	}
	// Copy in order not to chance values
	nums := make([]float64, len(values))
	copy(nums, values)
	sort.Float64s(nums)
	i := len(nums) / 2
	if len(nums)%2 == 1 {
		return nums[i], nil
	}
	v := (nums[i-1] + nums[i]) / 2
	return v, nil
}

func concat(s1, s2 []string) []string {
	// Restriction: no "for" loops
	if cap(s1) >= len(s1)+len(s2) {
		s1 = s1[:len(s1)+len(s2)]
		return s1
	} else {
		s3 := make([]string, len(s1)+len(s2))
		copy(s3, s1)
		copy(s3[len(s1):], s2)
		return s3
		//return append(s1, s2...)
	}
}

func appendInt(s []int, v int) []int {
	i := len(s)
	if len(s) < cap(s) { // enough space in underlying array to append
		s = s[:len(s)+1]
	} else { // need to reallocate and copy
		fmt.Printf("reallocating: %d ->%d\n", len(s), 2*len(s)+1)
		s2 := make([]int, 2*len(s)+1)
		copy(s2, s) // built-in function, remember that the parameter order is dest, src rather than src, dest
		s = s2[:len(s)+1]
	}
	s[i] = v
	return s
}

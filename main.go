package main

import (
	"fmt"
	"time"
)

func main() {
	// Uncomfortable code
	number, boolean := "Go Expert", 1
	fmt.Println("This is uncomfortable code.")
	fmt.Println("If you declare a variable called int, then you cannot use int as a type when you declare variables somewhere any more.")
	fmt.Println(number, boolean)
	fmt.Println("-------------------------")

	// struct
	var empty struct{}
	var point struct {
		ID   string
		x, y int
	}
	fmt.Println("Just want to resolve errors created by lint", empty, point)
	fmt.Println("-------------------------")

	var array [5]int
	arrayLiteral := [5]int{1, 2, 3, 4, 5}
	arrayInference := [...]int{1, 2, 3, 4, 5}
	arrayIndex := [...]int{2: 5, 5: 7, 7: 11}
	fmt.Println("array =: ", array)
	fmt.Println("arrayLiteral =: ", arrayLiteral)
	fmt.Println("arrayInference =: ", arrayInference)
	fmt.Println("arrayIndex =: ", arrayIndex)
	fmt.Println("-------------------------")

	var slice []int
	sliceLiteral := []int{1, 2, 3, 4, 5}
	fmt.Println("slice: ", slice)
	fmt.Println("sliceLiteral: ", sliceLiteral)
	fmt.Println("-------------------------")

	var m map[string]int
	mapLiteral := map[string]int{
		"Masataka": 185,
		"Hirotaka": 190,
	}
	fmt.Println("m: ", m)
	fmt.Println("mapLiteral: ", mapLiteral)
	fmt.Println("-------------------------")

	type MyDuration time.Duration
	d := MyDuration(100)
	fmt.Printf("The type of d is %T\n", d)

	td := time.Duration(d)
	fmt.Printf("The type of td is %T\n", td)

	md := 100 * d
	fmt.Printf("The type of md is %T\n", md)
}

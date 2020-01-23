package main

import "fmt"

func main() {
	// Arrays
	var fruitArr [2]string

	// Assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	// Declare and assign

	fruitArr2 := [2]string { "Apple", "Orange" }

	fmt.Println(fruitArr)
	fmt.Println(fruitArr2)


	// Slices
	fruitSlice := []string { "Apple", "Orange", "Grape", "Cherry" }
	fmt.Println(fruitSlice)
	fmt.Println(fruitSlice[1:2]) // Orange
	fmt.Println(fruitSlice[1:3]) // Orange, Grape
}
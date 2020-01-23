package main

import "fmt"

func main() {
	a := 5
	ptr := &a

	fmt.Println(a, *ptr)
	
	a = 10
	fmt.Println(a, *ptr)

	*ptr = 15
	fmt.Println(a, *ptr)

}

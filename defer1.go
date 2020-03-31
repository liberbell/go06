package main

import "fmt"

func main() {
	fmt.Println("\nFirst in ---> First out\n")
	for x := 0; x < 5; x++ {
		defer fmt.Println(x, "popped")
		fmt.Println(x, "deferred")
	}
}

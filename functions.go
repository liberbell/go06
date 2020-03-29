package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func subt(x int, y int) int {
	return x - y
}

func mult(a, b, string) (string, string) {
	return a, b
}

// func main() {
// 	fmt.Println(subt(18, 10))
// }

func main() {
	x, y := mult("hello", "world")
	fmt.Println(x, y)
}

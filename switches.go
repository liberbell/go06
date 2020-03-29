package main

import "fmt"

func main() {
	s := 5
	fmt.Println("This is: ", s)
	switch s {
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	case 5:
		fmt.Println("five")
	}
}

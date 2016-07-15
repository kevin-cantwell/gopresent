package main

import "fmt"

func main() {
	hello()
}

// START OMIT
// Now you see me
// END OMIT
// Now you don't
func hello() {
	fmt.Println("Hello, Presenters!")
}

package main

import (
	"fmt"
	"os"
)

func main() {

	name := ""

	for _, word := range os.Args[1:] {

		name = fmt.Sprintf("%v %v", name, word)

	}

	fmt.Println("Hello", name, "Welcome to the jungle")
}
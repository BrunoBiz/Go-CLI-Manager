package main

import (
	"fmt"
	"os"
)

var firstParam string

func main() {
	if len(os.Args) > 1 {
		firstParam = os.Args[1]
	} else {
		firstParam = "--"
	}

	fmt.Println(firstParam)
}

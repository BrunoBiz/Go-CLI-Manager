package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Print("MST - MAIN")

	if len(os.Args) > 1 {
		firstParam := os.Args[1]

		if firstParam == "start" {
			fmt.Println("MST - STARTED")
		}
	} else {
		fmt.Println("MST - ERR") // TODO - switch Log.fatal to errors
		return
	}
}

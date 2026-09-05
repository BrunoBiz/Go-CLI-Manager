package main

import (
	"fmt"
	"time"
)

func main() {
	var input string

	fmt.Println("Mock server started at: ", time.Now())

	for ok := true; ok; ok = (input != "") {
		n, err := fmt.Scanln(&input)

		if n < 1 || err != nil {
			fmt.Println("MST - INVALID INPUT")
			continue
		}

		if input == "shutdown" {
			fmt.Println("MST - SERVER STOPPING")
			time.Sleep(15 * time.Second)
			break
		}
	}
}

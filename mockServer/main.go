package main

import (
	"fmt"
	"time"
)

// IT STARTS WHEN CALLED, DOES NOT NEED A START PARAMETER
// WHEN CALLED, HAS TO BEGIN THE LOOP AND WAIT FOR THE SHUTDOWN COMMAND

func main() {
	var input string

	// REFACTOR ALL OF THIS, IT WILL START AND SIMPLY WAIT FOR THE SHUTDOWN COMMAND

	for ok := true; ok; ok = (input == "") {
		n, err := fmt.Scanln(&input)

		if n < 1 || err != nil {
			fmt.Println("MST - INVALID INPUT")
			continue
		}

		//fmt.Println(option)

		if input == "shutdown" {
			fmt.Println("MST - SERVER STOPPING")
			time.Sleep(3 * time.Second)
			break
		}
	}
}

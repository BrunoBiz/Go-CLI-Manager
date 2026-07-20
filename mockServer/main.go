package main

import (
	"fmt"
	"regexp"
	"time"
)

// IT STARTS WHEN CALLED, DOES NOT NEED A START PARAMETER
// WHEN CALLED, HAS TO BEGIN THE LOOP AND WAIT FOR THE SHUTDOWN COMMAND

func main() {
	fmt.Println("MST - MAIN")

	//	var serverOn bool
	var re = regexp.MustCompile(`shutdown`)
	var input string

	// REFACTOR ALL OF THIS, IT WILL START AND SIMPLY WAIT FOR THE SHUTDOWN COMMAND

	for ok := true; ok; ok = (re.MatchString(input)) {
		n, err := fmt.Scanln(&input)

		if n < 1 || err != nil {
			fmt.Println("MST - INVALID INPUT")
			continue
		}

		//fmt.Println(option)

		if input == "stop" {
			fmt.Println("MST - SERVER STOPPING")
			time.Sleep(3 * time.Second)
			break
		}

		if input == "restart" {
			fmt.Println("MST - SERVER RESTARTING...")
			time.Sleep(3 * time.Second)
			continue
		}

		if input == "status" {
			fmt.Println("MST - SERVER RUNNING")
			continue
		}
	}
}

package main

import (
	"fmt"
	"os"
	"regexp"
	"time"
)

func main() {
	fmt.Println("MST - MAIN")
	//	var serverOn bool
	var re = regexp.MustCompile(`start|stop|restart|status`)

	if len(os.Args) > 1 {
		option := os.Args[1]
		var input string

		switch option {
		case "start", "restart":
			if option == "restart" {
				fmt.Println("MST - SERVER RESTARTING")
				time.Sleep(3 * time.Second)
			}

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

		case "stop":
			fmt.Println("MST - SERVER NOT RUNNING")

		case "status":
			fmt.Println("MST - SERVER OFFLINE")

		}

	} else {
		fmt.Println("MST - ERR") // TODO - switch Log.fatal to errors
		return
	}
}

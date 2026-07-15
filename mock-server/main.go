package mockserver

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello there")

	if len(os.Args) > 1 {
		firstParam := os.Args[1]

		if firstParam == "start" {
			fmt.Println("START")
		}
	} else {
		fmt.Println("ERR") // TODO - switch Log.fatal to errors
		return
	}
}

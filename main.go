package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

var firstParam string

func main() {
	if len(os.Args) > 1 {
		firstParam = os.Args[1]
	} else {
		fmt.Println("ERR")
		return
	}

	switch firstParam {
	case "start":
		//fmt.Println("COMEÇOU")
		cmd := exec.Command("./ServerStart.sh")
		cmd.Dir = "/opt/minecraft/MBC-Server/"

		// 2. Connect the command's outputs directly to your terminal
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		fmt.Println(cmd.Args)

		// 3. Execute the script and wait for it to finish
		err := cmd.Run()
		if err != nil {
			log.Fatalf("Script failed to run: %v", err)
		}

		fmt.Println("Script executed successfully!")

	case "stop":
		fmt.Println("PAROU")
	case "restart":
		fmt.Println("RESTART")
	case "status":
		fmt.Println("STATUS")
	}

	fmt.Println(firstParam)
}

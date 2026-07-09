package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
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
		// Starts a new TMUX session running the server start shell
		cmd := exec.Command("tmux", "new", "-d", "-s", "minecraft_tmux", "'/opt/minecraft/MBC-Server/ServerStart.sh'")
		cmd.Dir = "/opt/minecraft/MBC-Server/"

		/*cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr*/

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
		// First checks if there is any tmux session running
		//cmd := exec.Command("tmux", "ls", "|", "grep", "'minecraft'", "-c")
		test, err := exec.Command("tmux", "ls").Output()

		/*err = cmd.Run()*/
		if err != nil {
			log.Fatalf("STATUS - Script failed to run: %v", err)
		}

		/*cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr*/

		fmt.Println(string(test))
		fmt.Println(strings.Contains(string(test), "minecraft_tmux"))

	}

	//fmt.Println(firstParam)
}

package main

import (
	gameservermgr "example/Go-CLI-Manager/gameServerMgr"
	"example/Go-CLI-Manager/util"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

var firstParam string

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load from config: ", err) // TODO - switch Log.fatal to errors
	}

	if len(os.Args) > 1 {
		firstParam = os.Args[1]
	} else {
		fmt.Println("ERR") // TODO - switch Log.fatal to errors
		return
	}

	gameServer := gameservermgr.NewGameServer(config)
	gameServer.OptionSwitch(firstParam)

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
		// First checks if the tmux session exists
		tmux_sessions, err := exec.Command("tmux", "ls").Output()

		if err != nil {
			log.Fatalf("STATUS - tmux ls - Script failed to run: %v", err)
		}

		if strings.Contains(string(tmux_sessions), "minecraft_tmux") {
			// tmux session exists
			fmt.Println("STATUS - TMUX Session found")

			// Sends the /spark tps command to check if the server is on
			_, err := exec.Command("tmux", "send-keys", "-t", "minecraft_tmux", "/spark tps", "ENTER").Output()

			if err != nil {
				log.Fatalf("STATUS - send-keys - Script failed to run: %v", err)
			}

			// Grabs the current panel from the tmux session
			tmux_pane, err := exec.Command("tmux", "capture-pane", "-p", "-t", "minecraft_tmux").Output()
			if err != nil {
				log.Fatalf("STATUS - capture-pane - Script failed to run: %v", err)
			}

			// Checks for the command output in the captured pane
			if strings.Contains(string(tmux_pane), "[?] TPS from last 5s, 10s, 1m, 5m, 15m:") {
				fmt.Println("STATUS - Server is running!")
			} else {
				log.Fatalf("STATUS - Server is offline!")
			}
		} else {
			// tmux does not exists
			log.Fatalf("STATUS - No TMUX Session found")
		}

	}
	//fmt.Println(firstParam)
}

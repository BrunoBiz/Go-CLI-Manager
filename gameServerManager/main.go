package main

import (
	gameservermgr "example/Go-CLI-Manager/gameServerManager/gameServerMgr"
	"example/Go-CLI-Manager/gameServerManager/util"
	"fmt"
	"log"
	"os"
)

func main() {
	config, err := util.LoadConfig("/opt/")
	if err != nil {
		log.Fatal("cannot load from config: ", err) // TODO - switch Log.fatal to errors
	}

	if len(os.Args) > 1 {
		gameServer := gameservermgr.NewGameServer(config)
		returnValue := gameServer.OptionSwitch(os.Args[1])

		returnValue.PrintError()
		returnValue.PrintMessage()
	} else {
		fmt.Println("ERR") // TODO - switch Log.fatal to errors
		return
	}
}

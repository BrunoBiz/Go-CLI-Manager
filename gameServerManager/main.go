package main

import (
	gameservermgr "example/Go-CLI-Manager/gameServerManager/gameServerMgr"
	"example/Go-CLI-Manager/gameServerManager/util"
	"fmt"
	"log"
	"os"
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
	returnValue := gameServer.OptionSwitch(firstParam)

	returnValue.PrintError()
	returnValue.PrintMessage()

}

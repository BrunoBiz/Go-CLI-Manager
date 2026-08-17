package main

import (
	gameservermgr "example/Go-CLI-Manager/gameServerManager/gameServerMgr"
	"example/Go-CLI-Manager/gameServerManager/logger"
	"example/Go-CLI-Manager/gameServerManager/util"
	"log/slog"
	"os"
)

func main() {
	config, err := util.LoadConfig("/opt/")
	if err != nil {
		slog.Error("ERR - Cannot load from config: " + err.Error())
		return
	}

	// Create/Load the log file and set as the default logger for SLOG
	err = logger.LoadLogger(config)
	if err != nil {
		slog.Error("ERR - Cannot load/create log file: " + err.Error())
		return
	}

	if len(os.Args) > 1 {
		gameServer := gameservermgr.NewGameServer(config)
		gameServer.OptionSwitch(os.Args[1], true)
	} else {
		slog.Error("Main -> empty os.Args")
		return
	}
}

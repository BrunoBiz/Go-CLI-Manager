package main

import (
	gameservermgr "example/Go-CLI-Manager/gameServerManager/gameServerMgr"
	"example/Go-CLI-Manager/gameServerManager/util"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"time"
)

func main() {
	config, err := util.LoadConfig("/opt/")
	if err != nil {
		slog.Error("ERR - Cannot load from config: " + err.Error())
		return
	}

	// Create/Load the log file and set as the default logger for SLOG
	err = LoadLogger(config)
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

func LoadLogger(config util.Config) error {
	//The 'Log' folder will always be in the root directory of GameServerManager
	err := checkDirectory()
	if err != nil {
		return err
	}

	// Every log file will be named according to the current date
	fileName := time.Now().Format(time.DateOnly)
	logFile, err := os.OpenFile(fmt.Sprintf("Log/%s-gsm.log", fileName), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		slog.Error("Can't create log file - " + err.Error())
		return err
	}

	logger := slog.New(slog.NewJSONHandler(logFile, nil))
	slog.SetDefault(logger)

	return nil
}

func checkDirectory() error {
	cmd := exec.Command(`/bin/bash`, `-c`, `if [ ! -d Log ]; then mkdir 'Log' && echo 'Created'; else echo 'Already exists'; fi`) // TODO - There might be a better way to do this, but it works
	tmux_sd, err := cmd.CombinedOutput()

	slog.Debug("LOGGER")
	slog.Debug(cmd.String())
	slog.Debug(string(tmux_sd))

	if err != nil {
		slog.Error("Can't create log directory - " + err.Error())
		return err
	}

	return nil
}

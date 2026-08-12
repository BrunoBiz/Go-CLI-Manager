package logger

import (
	"example/Go-CLI-Manager/gameServerManager/util"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"time"
)

func NewLogger(config util.Config) bool {
	if checkDirectory() != nil {
		return false
	}

	fileName := time.Now().Format(time.DateOnly)
	logFile, err := os.OpenFile(fmt.Sprintf("Log/%s-gsm.log", fileName), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		slog.Error("Can't create log file - " + err.Error())
		return false
	}

	logHandler := slog.NewJSONHandler(logFile, nil)
	logger := slog.New(logHandler)
	slog.SetDefault(logger)

	return true
}

func checkDirectory() error {
	cmd := exec.Command(`/bin/bash`, `-c`, `if [ ! -d Log ]; then mkdir 'Log' && echo 'Created'; else echo 'Already exists'; fi`) // TODO - There might be a better way to do this, but it works

	tmux_sd, err := cmd.CombinedOutput()

	slog.Debug("LOGGER")
	slog.Debug(cmd.String())
	slog.Info(string(tmux_sd))

	if err != nil {
		slog.Error("Can't create log directory - " + err.Error())
		return err
	}

	return nil
}

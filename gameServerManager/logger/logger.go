package logger

import (
	"example/Go-CLI-Manager/gameServerManager/util"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/exec"
	"time"
)

func LoadLogger(config util.Config) error {
	//The 'Log' folder will always be in the root directory of GameServerManager
	err := checkDirectory()
	if err != nil {
		return err
	}

	// Will delete any file older than 15 days
	err = logRotation()
	if err != nil {
		return err
	}

	// Every log file will be named accordingly to the current date
	fileName := time.Now().Format(time.DateOnly)
	logFile, err := os.OpenFile(fmt.Sprintf("Log/%s-gsm.log", fileName), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		slog.Error("Can't create log file - " + err.Error())
		return err
	}

	mWriter := io.MultiWriter(logFile, os.Stdout)

	logger := slog.New(slog.NewTextHandler(mWriter, &slog.HandlerOptions{
		//AddSource: true, // Adds source=main.go:15 to the log line
	}))
	slog.SetDefault(logger)

	return nil
}

func checkDirectory() error {
	cmd := exec.Command(`/bin/bash`, `-c`, `if [ ! -d Log ]; then mkdir 'Log' && echo 'Created'; else echo 'Already exists'; fi`) // TODO - There might be a better way to do this, but it works
	checkDir, err := cmd.CombinedOutput()

	slog.Debug("checkDirectory()")
	slog.Debug(cmd.String())
	slog.Debug(string(checkDir))

	// The return here (checkDir) does not matter, either it returns `Created` or  `Already exists`, both scenarios are acceptable, all that matters is that there was no errors
	if err != nil {
		slog.Error("Can't create log directory - " + err.Error())
		return err
	}

	return nil
}

func logRotation() error {
	// Will delete any log file older than 15 days
	cmd := exec.Command(`/bin/bash`, `-c`, `find`, `./Log`, `-type`, `f`, `-mtime`, `+15`, `-exec`, `rm`, `{}`, `';'`)
	deleteOldLogs, err := cmd.CombinedOutput()

	slog.Debug("logRotation()")
	slog.Debug(cmd.String())
	slog.Debug(string(deleteOldLogs))

	if err != nil {
		slog.Error("Can't delete old logs - " + err.Error())
		return err
	}

	return nil
}

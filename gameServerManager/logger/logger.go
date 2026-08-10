package logger

import (
	"example/Go-CLI-Manager/gameServerManager/util"
	"log/slog"
	"os/exec"
)

type Logger struct {
	logDir string
}

func NewLogger(config util.Config) *Logger {
	logDir := config.GameServerDir + "/Log"

	checkDirectory()

	logger := &Logger{
		logDir: logDir,
	}

	return logger
}

func checkDirectory() (bool, error) {
	// if [ ! -d ./Log ]; then mkdir "Log" && echo "Created"; else echo "Already exists"; fi

	//cmd := exec.Command(`/bin/bash`, `-c`, `"if [ ! -d /Log ]; then mkdir 'Log' && echo 'Created'; else echo 'Already exists'; fi"`)
	cmd := exec.Command(`/bin/bash`, `-c`, `echo 'Hello from Go'`)

	/*`if`,
	`[`,
	`!`,
	`-d`,
	`./Log`,
	`]`,
	`;`,
	`then`,
	`mkdir`,
	`'Log'`,
	`&&`,
	`echo`,
	`'Created'`,
	`;`,
	`else`,
	`echo`,
	`'Already exists'`,
	`;`,
	`fi`,
	`"`)*/

	tmux_sd, err := cmd.CombinedOutput()
	slog.Info("LOGGER")
	slog.Info(cmd.String())
	slog.Info(string(tmux_sd))

	if err != nil {
		slog.Error(err.Error())
		return false, err
	}

	return false, nil
}

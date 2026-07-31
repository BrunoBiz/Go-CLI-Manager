package gameservermgr

import (
	"fmt"
	"log/slog"
	"strconv"
	"strings"
)

type ReturnValue struct {
	option        string
	command       string
	commandResult string
	success       bool
	serverOnline  bool // Only used by the status option -> stores if the server is On
	message       string
	err           error
}

func (returnValue *ReturnValue) PrintLogs() {
	slog.Info("GSM - " + strings.ToUpper(returnValue.option) + " - MESSAGE: " + returnValue.message)
	slog.Info("GSM - " + strings.ToUpper(returnValue.option) + " - SUCCESS: " + strconv.FormatBool(returnValue.success))
	if returnValue.option == "status" {
		slog.Info("GSM - " + strings.ToUpper(returnValue.option) + " - SERVER STATUS: " + strconv.FormatBool(returnValue.serverOnline))
	}

	if returnValue.err != nil {
		slog.Error("GSM - " + strings.ToUpper(returnValue.option) + " - ERROR: " + returnValue.err.Error())
	}

	if returnValue.command != "" {
		slog.Info("GSM - " + strings.ToUpper(returnValue.option) + " - COMMAND: " + returnValue.command)
	}

	if returnValue.commandResult != "" {
		slog.Info("GSM - " + strings.ToUpper(returnValue.option) + " - COMMAND RESULT: " + returnValue.commandResult)
	}
	fmt.Println()
}

func newReturnValue(option string, command string, commandResult string, success bool, serverOnline bool, message string, err error) ReturnValue {
	return ReturnValue{
		option:        option,
		command:       command,
		commandResult: commandResult,
		success:       success,
		serverOnline:  serverOnline,
		message:       message,
		err:           err,
	}
}

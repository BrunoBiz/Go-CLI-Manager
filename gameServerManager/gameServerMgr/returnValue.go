package gameservermgr

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"strconv"
	"strings"
)

type ReturnValue struct {
	Option        string
	Command       string
	CommandResult string
	Success       bool
	ServerOnline  bool // Only used by the details option -> stores if the server is On
	Message       string
	ErrMsg        string
}

func (returnValue *ReturnValue) PrintLogs() {
	slog.Info("GSM - " + strings.ToUpper(returnValue.Option) + " - SUCCESS: " + strconv.FormatBool(returnValue.Success))
	if returnValue.Option == "details" {
		var serverStatus string
		if returnValue.ServerOnline {
			serverStatus = "STARTED"
		} else {
			serverStatus = "STOPPED"
		}

		slog.Info("GSM - " + strings.ToUpper(returnValue.Option) + " - SERVER STATUS: " + serverStatus)
	}

	if returnValue.ErrMsg != "" {
		slog.Error("GSM - " + strings.ToUpper(returnValue.Option) + " - ERROR: " + returnValue.ErrMsg)
	}

	if returnValue.Command != "" {
		slog.Info("GSM - " + strings.ToUpper(returnValue.Option) + " - COMMAND: " + returnValue.Command)
	}

	if returnValue.CommandResult != "" {
		slog.Info("GSM - " + strings.ToUpper(returnValue.Option) + " - COMMAND RESULT: " + returnValue.CommandResult)
	}
	slog.Info("GSM - " + strings.ToUpper(returnValue.Option) + " - MESSAGE: " + returnValue.Message)
}

func (returnValue *ReturnValue) PrintLogsJSON() {

	// Currently not used - Opted for a more streamlined option, more compatible with LinuxGSM

	jsonFormattedLog, err := json.Marshal(returnValue)
	if err != nil {
		fmt.Println("err " + err.Error())
	}
	fmt.Println(string(jsonFormattedLog)) // -> StdOut
	slog.Info(string(jsonFormattedLog))   // -> Log
}

func newReturnValue(option string, command string, commandResult string, success bool, serverOnline bool, message string, err error) ReturnValue {
	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}

	return ReturnValue{
		Option:        option,
		Command:       command,
		CommandResult: commandResult,
		Success:       success,
		ServerOnline:  serverOnline,
		Message:       message,
		ErrMsg:        errMsg,
	}
}

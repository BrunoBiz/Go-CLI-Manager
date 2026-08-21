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
	slog.Info("GSM - " + strings.ToUpper(returnValue.Option) + " - MESSAGE: " + returnValue.Message)
	slog.Info("GSM - " + strings.ToUpper(returnValue.Option) + " - SUCCESS: " + strconv.FormatBool(returnValue.Success))
	if returnValue.Option == "details" {
		slog.Info("GSM - " + strings.ToUpper(returnValue.Option) + " - SERVER STATUS: " + strconv.FormatBool(returnValue.ServerOnline))
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
	fmt.Println()
}

func (returnValue *ReturnValue) PrintLogsJSON() {
	/*var errMsg string
	if returnValue.err != nil {
		errMsg = returnValue.err.Error()
	}*/

	/*jsonFormattedLog := fmt.Sprintf(
		`{
			"option": %q,
			"message": %q,
			"success": %t,
			"status": %t,
			"error": %q,
			"command": %q,
			"commandResult": %q
		}`,
		returnValue.option,
		returnValue.message,
		returnValue.success,
		returnValue.serverOnline,
		errMsg,
		returnValue.command,
		strings.TrimSpace(returnValue.commandResult),
	)*/
	//slog.Info(jsonFormattedLog)

	jsonFormattedLog, err := json.Marshal(returnValue)
	if err != nil {
		fmt.Println("err " + err.Error())
	}
	fmt.Println(string(jsonFormattedLog))
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

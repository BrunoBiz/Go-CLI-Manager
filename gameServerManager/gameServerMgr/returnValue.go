package gameservermgr

import "log/slog"

type ReturnValue struct {
	option        string
	command       string
	commandResult string
	success       bool
	message       string
	err           error
}

func (returnValue *ReturnValue) PrintMessage() {
	slog.Info("GSM - MESSAGE: " + returnValue.message)
}

func (returnValue *ReturnValue) PrintError() {
	if returnValue.err != nil {
		slog.Error("GSM - ERROR: " + returnValue.err.Error())
	}
}

func (returnValue *ReturnValue) PrintCommand() {
	if returnValue.command != "" {
		slog.Info("GSM - COMMAND: " + returnValue.command)
	}
}

func (returnValue *ReturnValue) PrintCommandResult() {
	if returnValue.commandResult != "" {
		slog.Info("GSM - COMMAND RESULT: " + returnValue.commandResult)
	}
}

func newReturnValue(option string, command string, commandResult string, success bool, message string, err error) ReturnValue {
	return ReturnValue{
		option:        option,
		command:       command,
		commandResult: commandResult,
		success:       success,
		message:       message,
		err:           err,
	}
}

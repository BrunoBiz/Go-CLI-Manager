package gameservermgr

import (
	"errors"
	"example/Go-CLI-Manager/gameServerManager/util"
	"log/slog"
)

type GameServer struct {
	tmuxSessionName string
	startFilePath   string
	gameServerDir   string
}

type ReturnValue struct {
	option  string
	valor   int
	ativo   bool
	message string
	err     error
}

func (returnValue *ReturnValue) PrintMessage() {
	slog.Info("GSM - MESSAGE: " + returnValue.message)
}

func (returnValue *ReturnValue) PrintError() {
	if returnValue.err != nil {
		slog.Error("GSM - ERROR: " + returnValue.err.Error())
	}
}

func newReturnValue(option string, valor int, ativo bool, message string, err error) ReturnValue {
	return ReturnValue{
		option:  option,
		valor:   valor,
		ativo:   ativo,
		message: message,
		err:     err,
	}
}

func NewGameServer(config util.Config) *GameServer {
	gameServer := &GameServer{
		tmuxSessionName: config.TMUXSessionName,
		startFilePath:   config.GameStartFilePath,
		gameServerDir:   config.GameServerDir,
	}

	return gameServer
}

func (gameServer *GameServer) OptionSwitch(option string) ReturnValue {
	switch option {
	case "start":
		return gameServer.start()
	case "stop":
		return gameServer.stop()
	case "restart":
		return gameServer.restart()
	case "status":
		return gameServer.status()
	}

	return newReturnValue("INVALID", 0, false, "OPCAO INVALIDA/NULA/VAZIA", errors.New("OPCAO INVALIDA/NULA/VAZIA\rOPCAO: "+option))
}

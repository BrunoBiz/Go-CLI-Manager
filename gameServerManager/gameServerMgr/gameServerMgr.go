package gameservermgr

import (
	"errors"
	"example/Go-CLI-Manager/gameServerManager/util"
)

type GameServer struct {
	tmuxSessionName string
	startFilePath   string
	gameServerDir   string
}

func NewGameServer(config util.Config) *GameServer {
	gameServer := &GameServer{
		tmuxSessionName: config.TMUXSessionName,
		startFilePath:   config.GameStartFilePath,
		gameServerDir:   config.GameServerDir,
	}

	return gameServer
}

func (gameServer *GameServer) OptionSwitch(option string, printLogs bool) ReturnValue {
	var returnSwitch ReturnValue

	switch option {
	case "start":
		returnSwitch = gameServer.start()
	case "stop":
		returnSwitch = gameServer.stop()
	case "restart":
		returnSwitch = gameServer.restart()
	case "status":
		returnSwitch = gameServer.status()
	default:
		returnSwitch = newReturnValue("INVALID", option, "", false, "INVALID/NIL/EMPTY OPTION", errors.New("OPCAO INVALIDA/NULA/VAZIA\rOPCAO: "+option))
	}

	if printLogs {
		returnSwitch.PrintLogs()
	}

	return returnSwitch
}

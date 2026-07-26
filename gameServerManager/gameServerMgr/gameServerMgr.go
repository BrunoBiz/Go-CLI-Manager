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

	return newReturnValue("INVALID", option, "", false, "INVALID/NIL/EMPTY OPTION", errors.New("OPCAO INVALIDA/NULA/VAZIA\rOPCAO: "+option))
}

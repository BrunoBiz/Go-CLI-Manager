package gameservermgr

import (
	"errors"
	"example/Go-CLI-Manager/util"
	"os/exec"
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

	return newReturnValue("INVALIDO", 0, false, "OPCAO INVALIDA/NULA/VAZIA", errors.New("OPCAO INVALIDA/NULA/VAZIA"))
}

func (gameServer *GameServer) start() ReturnValue {
	// Starts a new TMUX session running the server start shell
	cmd := exec.Command("tmux", "new", "-d", "-s", gameServer.tmuxSessionName, gameServer.startFilePath)
	cmd.Dir = gameServer.gameServerDir
	err := cmd.Run()

	if err != nil {
		return newReturnValue("start", 0, false, "FAILED TO START", err)
	} else {
		return newReturnValue("start", 0, true, "STARTED", err)
	}
}

func (gameServer *GameServer) stop() ReturnValue {
	return newReturnValue("", 0, false, "", nil)
}

func (gameServer *GameServer) restart() ReturnValue {

	return newReturnValue("", 0, false, "", nil)
}

func (gameServer *GameServer) status() ReturnValue {
	return newReturnValue("", 0, false, "", nil)
}

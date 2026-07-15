package gameservermgr

import (
	"errors"
	"example/Go-CLI-Manager/gameServerManager/util"
	"os/exec"
	"strings"
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
	Message string
	err     error
}

func newReturnValue(option string, valor int, ativo bool, message string, err error) ReturnValue {
	return ReturnValue{
		option:  option,
		valor:   valor,
		ativo:   ativo,
		Message: message, // TODO - Maybe change back to private
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

	return newReturnValue("INVALIDO", 0, false, "OPCAO INVALIDA/NULA/VAZIA", errors.New("OPCAO INVALIDA/NULA/VAZIA\rOPCAO: "+option))
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
	/*
		NOT a health check, this only validates that the TMUX session is active, the game may have crashed,
		is frozen or an error may have ocurred and it will still return as ONLINE if the session is still on!
	*/

	// First checks if the tmux session exists
	tmux_session, err := exec.Command("tmux", "ls").Output()

	if err != nil {
		return newReturnValue("status", 0, false, "tmux ls - Script failed to run", err)
	}

	if strings.Contains(string(tmux_session), "minecraft_tmux") {
		// tmux session exists
		return newReturnValue("status", 1, true, "Server running", nil)
	} else {
		// tmux does not exists
		return newReturnValue("status", 0, false, "Server offline", nil)
	}
}

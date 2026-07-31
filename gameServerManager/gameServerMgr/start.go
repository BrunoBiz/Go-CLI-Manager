package gameservermgr

import (
	"os/exec"
)

func (gameServer *GameServer) start() ReturnValue {
	// Starts a new TMUX session running the server start shell
	cmd := exec.Command("tmux", "new", "-d", "-s", gameServer.tmuxSessionName, gameServer.startFilePath)
	cmd.Dir = gameServer.gameServerDir
	tmux_start, err := cmd.CombinedOutput()

	if err != nil { // TODO - Refactor -> remove IF
		return newReturnValue("start", cmd.String(), string(tmux_start), false, "FAILED TO START", err)
	} else {
		return newReturnValue("start", cmd.String(), string(tmux_start), false, "STARTED", err)
	}
}

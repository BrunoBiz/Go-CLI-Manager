package gameservermgr

import (
	"log/slog"
	"os/exec"
)

func (gameServer *GameServer) start() ReturnValue {
	// Starts a new TMUX session running the server start shell
	cmd := exec.Command("tmux", "new", "-d", "-s", gameServer.tmuxSessionName, gameServer.startFilePath)
	cmd.Dir = gameServer.gameServerDir
	tmux_start, err := cmd.Output()

	if err != nil { // TODO - Refactor -> remove IF
		slog.Info("FAILED TO START")
		return newReturnValue("start", cmd.String(), string(tmux_start), false, "FAILED TO START", err)
	} else {
		slog.Info("STARTED")
		return newReturnValue("start", cmd.String(), string(tmux_start), true, "STARTED", err)
	}
}

package gameservermgr

import (
	"log/slog"
	"os/exec"
)

func (gameServer *GameServer) start() ReturnValue {
	// Starts a new TMUX session running the server start shell
	cmd := exec.Command("tmux", "new", "-d", "-s", gameServer.tmuxSessionName, gameServer.startFilePath)
	cmd.Dir = gameServer.gameServerDir
	err := cmd.Run()

	slog.Info("Comando executado: " + cmd.String())
	slog.Info("Diretorio: " + cmd.Dir)

	if err != nil {
		slog.Info("FAILED TO START")
		return newReturnValue("start", 0, false, "FAILED TO START", err)
	} else {
		slog.Info("STARTED")
		return newReturnValue("start", 0, true, "STARTED", err)
	}
}

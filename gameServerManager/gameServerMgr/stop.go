package gameservermgr

import (
	"log/slog"
	"os/exec"
)

func (gameServer *GameServer) stop() ReturnValue {
	cmd := exec.Command("tmux", "send-keys", "-t", gameServer.tmuxSessionName, "shutdown", "ENTER")
	tmux_sd, err := cmd.Output()

	slog.Info("Comando executado: " + cmd.String())
	slog.Info("Retorno comando: " + string(tmux_sd))

	if err != nil {
		return newReturnValue("stop", 0, false, "shutdown - Script failed to run", err)
	}

	return newReturnValue("stop", 1, false, "STOPED", nil)
}

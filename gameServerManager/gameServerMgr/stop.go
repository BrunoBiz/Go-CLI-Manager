package gameservermgr

import (
	"os/exec"
)

func (gameServer *GameServer) stop() ReturnValue {
	cmd := exec.Command("tmux", "send-keys", "-t", gameServer.tmuxSessionName, "shutdown", "ENTER")
	tmux_sd, err := cmd.CombinedOutput()

	if err != nil {
		return newReturnValue("stop", cmd.String(), string(tmux_sd), false, "shutdown - Script failed to run", err)
	}

	return newReturnValue("stop", cmd.String(), string(tmux_sd), false, "Stopping server...", nil)
}

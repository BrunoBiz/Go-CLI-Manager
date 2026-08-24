package gameservermgr

import (
	"os/exec"
	"strings"
)

func (gameServer *GameServer) start() ReturnValue {
	// Starts a new TMUX session running the server start shell
	cmd := exec.Command("tmux", "new", "-d", "-s", gameServer.config.TMUXSessionName, gameServer.config.GameStartFilePath)
	cmd.Dir = gameServer.config.GameServerDir
	tmux_start, err := cmd.CombinedOutput()

	if err != nil && strings.Contains(string(tmux_start), "duplicate session:") { // Server already running
		return newReturnValue("start", cmd.String(), string(tmux_start), false, true, "Server is already running", err)
	}

	if err != nil { // Any other error - Server failed to start
		return newReturnValue("start", cmd.String(), string(tmux_start), false, false, "Unable to start", err)
	}

	// Server started successfully
	return newReturnValue("start", cmd.String(), string(tmux_start), true, false, "Server started", err)
}

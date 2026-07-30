package gameservermgr

import (
	"fmt"
	"os/exec"
	"strings"
)

func (gameServer *GameServer) status2() ReturnValue {
	/*
		NOT a health check, this only validates that the TMUX session is active, the game may have crashed,
		is frozen or an error may have ocurred and it will still return as ONLINE if the session is still on!
	*/

	// First checks if the tmux session exists
	cmd := exec.Command("tmux", "ls")
	tmux_session, err := cmd.CombinedOutput()

	fmt.Println("-------------------")
	fmt.Println(string(tmux_session))
	fmt.Println("-------------------")

	if err != nil {
		return newReturnValue("status", cmd.String(), string(tmux_session), false, "tmux ls - Script failed to run", err)
	}

	if strings.Contains(string(tmux_session), gameServer.tmuxSessionName) {
		// tmux session exists
		return newReturnValue("status", cmd.String(), string(tmux_session), true, "Server running", nil)
	} else {
		// tmux does not exists
		return newReturnValue("status", cmd.String(), string(tmux_session), false, "Server offline", nil)
	}
}

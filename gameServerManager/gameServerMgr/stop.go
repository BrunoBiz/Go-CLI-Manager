package gameservermgr

import (
	"errors"
	"fmt"
	"log/slog"
	"os/exec"
	"strings"
	"time"
)

func (gameServer *GameServer) stop() ReturnValue {
	var stopServerDetails ReturnValue

	//exec.Command("echo", `"Stopping server..."`)
	slog.Info("Stopping server...")

	cmd := exec.Command("tmux", "send-keys", "-t", gameServer.config.TMUXSessionName, "shutdown", "ENTER")
	tmux_sd, err := cmd.CombinedOutput()

	// Server is already stopped
	if err != nil && strings.Contains(string(tmux_sd), "no server running on") {
		return newReturnValue("stop", cmd.String(), string(tmux_sd), false, false, "Server is already stopped", err)
	}

	// Any other error
	if err != nil {
		return newReturnValue("stop", cmd.String(), string(tmux_sd), false, false, "shutdown - Script failed to run", err)
	}

	// max timeout wait
	currentTime := time.Now()
	timeOut := currentTime.Add(time.Duration(gameServer.config.ServerStopTimeout) * time.Second)

	// Waits until the server is offline
	for {
		fmt.Printf("\rTime elapsed: %s / %ds", time.Since(currentTime).Round(time.Second), gameServer.config.ServerStopTimeout)

		stopServerDetails = gameServer.OptionSwitch("details", false)
		if !stopServerDetails.ServerOnline {
			// Server stopped
			fmt.Print("\n\n") // TODO - super ugly code
			break
		}

		// Time out - server did not shutdown in time
		if time.Now().After(timeOut) {
			return newReturnValue("stop", "", "", false, false, "Server timed out", errors.New("Timeout"))
		}

		time.Sleep(1 * time.Second) // TODO - might need to remove this
	}

	return newReturnValue("stop", cmd.String(), string(tmux_sd), true, false, "Server stopped", nil)
}

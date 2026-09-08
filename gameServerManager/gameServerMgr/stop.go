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
	var err error
	var cmd *exec.Cmd
	var tmux_sd []byte

	slog.Info("Stopping server...")

	// Some servers may require custom shutdown sequences
	/*
		CustomShutdownSequence
		0 - Default - Shutdown command
		1 - Shutdown command followed by a N(o) commnand
	*/
	if gameServer.config.CustomShutdownSequence == 0 {
		cmd = exec.Command("tmux", "send-keys", "-t", gameServer.config.TMUXSessionName, "shutdown", "ENTER")
		tmux_sd, err = cmd.CombinedOutput()
	}

	if gameServer.config.CustomShutdownSequence == 1 {
		/*
			After shutdown, server prompts user to restart automatically after 30s.
			Need to send a 'N' to stop the auto-restart and keep the server off.

		*/
		cmd = exec.Command("tmux", "send-keys", "-t", gameServer.config.TMUXSessionName, "shutdown", "ENTER")
		tmux_sd, err = cmd.CombinedOutput()

		// Will need to read the stdout and wait here until 'Server will re-start *automatically* in less than 30 seconds...' shows up
		// Other than that, all works fine

		cmd = exec.Command("tmux", "send-keys", "-t", gameServer.config.TMUXSessionName, "n", "ENTER")

	}

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

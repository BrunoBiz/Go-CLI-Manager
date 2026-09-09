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
	var cmd, cmdNCommand *exec.Cmd
	var tmuxStop, tmuxCapturePane []byte

	slog.Info("Stopping server...")

	// Some servers may require custom shutdown sequences
	/*
		CustomShutdownSequence
		0 - Default - Shutdown command
		1 - Shutdown command followed by a N(o) commnand
	*/
	if gameServer.config.CustomShutdownSequence == 0 {
		slog.Debug("Shutdown - Sequence 0")

		cmd = exec.Command("tmux", "send-keys", "-t", gameServer.config.TMUXSessionName, "shutdown", "ENTER")
		tmuxStop, err = cmd.CombinedOutput()
	}

	if gameServer.config.CustomShutdownSequence == 1 {
		slog.Debug("Shutdown - Sequence 1")

		/*
			After shutdown, server prompts user to restart automatically after 30s.
			Need to send a 'N' to stop the auto-restart and keep the server off.

		*/
		cmd = exec.Command("tmux", "send-keys", "-t", gameServer.config.TMUXSessionName, "shutdown", "ENTER")
		tmuxStop, err = cmd.CombinedOutput() // send-keys has no output, fire and forget

		// Will need to read the stdout and wait here until 'Server will re-start *automatically* in less than 30 seconds...' shows up
		for {
			// Captures the last 100 lines from the tmux terminal
			cmd = exec.Command("tmux", "capture-pane", "-S", "-100", "-E", "-", "-p", "-t", gameServer.config.TMUXSessionName)
			tmuxCapturePane, err = cmd.CombinedOutput()

			if err != nil {
				return newReturnValue("stop", cmd.String(), string(tmuxCapturePane), false, false, "shutdown - Script failed to run (0)", err)
			}

			if strings.Contains(string(tmuxCapturePane), "Server will re-start ") {
				cmdNCommand = exec.Command("tmux", "send-keys", "-t", gameServer.config.TMUXSessionName, "n", "ENTER")
				_, err = cmdNCommand.CombinedOutput()

				if err != nil {
					return newReturnValue("stop", cmdNCommand.String(), "", false, false, "shutdown - Script failed to run (1)", err)
				}

				break
			}

			time.Sleep(time.Second)
		}
	}

	// Server is already stopped
	if err != nil && strings.Contains(string(tmuxStop), "no server running on") {
		return newReturnValue("stop", cmd.String(), string(tmuxStop), false, false, "Server is already stopped", err)
	}

	// Any other error
	if err != nil {
		return newReturnValue("stop", cmd.String(), string(tmuxStop), false, false, "shutdown - Script failed to run (2)", err)
	}

	// max timeout wait
	slog.Debug("Max timeout")
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

		time.Sleep(time.Second) // TODO - might need to remove this
	}

	return newReturnValue("stop", cmd.String(), string(tmuxStop), true, false, "Server stopped", nil)
}

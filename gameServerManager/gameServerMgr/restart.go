package gameservermgr

import (
	"errors"
	"time"
)

func (gameServer *GameServer) restart() ReturnValue {
	// First checks if the server is running
	returnDetails := gameServer.OptionSwitch("details", true)
	var stopServerDetails ReturnValue

	// SERVER RUNNING -> Stops it
	if returnDetails.serverOnline {
		returnStop := gameServer.OptionSwitch("stop", true)

		if !returnStop.success { // Cant stop server
			return returnStop
		} else { // TODO - needs more testing
			// Waits until the server is offline
			timeOut := time.Now().Add(time.Duration(gameServer.config.ServerStopTimeout) * time.Second)
			for {
				stopServerDetails = gameServer.OptionSwitch("details", false)
				if !stopServerDetails.serverOnline {
					// Server stopped
					break
				}

				// Time out - server did not shutdown in time
				if time.Now().After(timeOut) {
					return newReturnValue("restart", "", "", false, false, "TIME OUT", errors.New("Timeout"))
				}

				time.Sleep(5 * time.Second)
			}
		}
	}

	// Start server
	returnStart := gameServer.OptionSwitch("start", false) // The last logs are already printed by the main restart function call
	return returnStart
}

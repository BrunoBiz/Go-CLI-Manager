package gameservermgr

import (
	"errors"
	"time"
)

func (gameServer *GameServer) restart() ReturnValue {
	// First checks if the server is running
	returnStatus := gameServer.OptionSwitch("status", true)
	var stopServerStatus ReturnValue

	// if so, stops it
	if returnStatus.serverOnline {
		returnStop := gameServer.OptionSwitch("stop", true)

		if !returnStop.serverOnline { // Cant stop server
			return returnStop
		} else { // TODO - needs more testing
			// Waits until the server is offline
			timeOut := time.Now().Add(time.Duration(gameServer.stopServerTimeOut) * time.Second)
			for {
				stopServerStatus = gameServer.OptionSwitch("status", false)
				if !stopServerStatus.serverOnline {
					break
				}

				// Time out - server did not shutdown after the set time
				if time.Now().After(timeOut) {
					return newReturnValue("restart", "", "", false, "TIME OUT", errors.New("Timeout"))
				}

				time.Sleep(5 * time.Second)
			}
		}
	}

	// starts server
	returnStart := gameServer.OptionSwitch("start", false) // The last logs are already printed by the main restart function call
	return returnStart
}

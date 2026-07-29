package gameservermgr

import (
	"time"
)

func (gameServer *GameServer) restart() ReturnValue {
	// First checks if the server is running
	returnStatus := gameServer.OptionSwitch("status", true)
	var stopServerStatus ReturnValue

	// if so, stops it
	if returnStatus.success {
		returnStop := gameServer.OptionSwitch("stop", true)

		if !returnStop.success { // Cant stop server
			return returnStop
		} else { // TODO - needs more testing
			// Waits until the server is offline
			timeOut := time.Now().Add(time.Duration(gameServer.stopServerTimeOut) * time.Second)
			for {
				stopServerStatus = gameServer.OptionSwitch("status", false)
				// "success" returns true or false depending on if the server is running or not when calling the status option;
				// When the return value is false, it means the server went offline, so we can proceed
				if !stopServerStatus.success {
					break
				}

				// Time out - server did not shutdown after the set time
				if time.Now().After(timeOut) {
					//return newReturnValue("restart", "", "", false, "TIME OUT", error.Error("Timeout"))
				}

				time.Sleep(5 * time.Second)
			}
		}
	}

	// starts server
	returnStart := gameServer.OptionSwitch("start", false) // The last logs are already printed by the main restart function call
	return returnStart
}

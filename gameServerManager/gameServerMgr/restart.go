package gameservermgr

func (gameServer *GameServer) restart() ReturnValue {
	// First checks if the server is running
	returnDetails := gameServer.OptionSwitch("details", true)

	// SERVER RUNNING -> Stops it
	if returnDetails.serverOnline {
		returnStop := gameServer.OptionSwitch("stop", true)

		if !returnStop.success { // Cant stop server
			return returnStop
		}
	}

	// Start server
	returnStart := gameServer.OptionSwitch("start", false) // The last logs are already printed by the main restart function call
	return returnStart
}

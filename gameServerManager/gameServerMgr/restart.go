package gameservermgr

func (gameServer *GameServer) restart() ReturnValue {
	// First checks if the server is running
	returnStatus := gameServer.OptionSwitch("status", true)

	// if so, stops it
	if returnStatus.success {
		returnStop := gameServer.OptionSwitch("stop", true)

		if !returnStop.success {
			return returnStop
		}
	}

	// starts server
	returnStart := gameServer.OptionSwitch("start", false) // The last logs are already printed by the main restart function call
	return returnStart
}

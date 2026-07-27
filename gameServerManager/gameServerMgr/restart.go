package gameservermgr

import (
	"fmt"
	"time"
)

func (gameServer *GameServer) restart() ReturnValue {
	// First checks if the server is running
	returnStatus := gameServer.OptionSwitch("status", true)

	// if so, stops it
	if returnStatus.success {
		returnStop := gameServer.OptionSwitch("stop", true)

		if !returnStop.success { // Cant stop server
			return returnStop
		} else { // TODO - needs more testing
			// Waits until the server is offline
			t2 := time.Now().Add(300 * time.Second)
			for {

				if !gameServer.OptionSwitch("status", true).success {
					fmt.Println("PARADO") // TODO - CLEAR
					break
				}

				if time.Now().After(t2) {
					fmt.Println("DEPOIS") // TODO - CLEAR
					break
				}

				time.Sleep(5 * time.Second)
			}
		}
	}

	// starts server
	returnStart := gameServer.OptionSwitch("start", false) // The last logs are already printed by the main restart function call
	return returnStart
}

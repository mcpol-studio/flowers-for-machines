package main

import (
	"time"

	"github.com/mcpol-studio/flowers-for-machines/client"
	"github.com/mcpol-studio/flowers-for-machines/game_control/game_interface"
	"github.com/mcpol-studio/flowers-for-machines/game_control/resources_control"

	"github.com/pterm/pterm"
)

var (
	c         *client.Client
	resources *resources_control.Resources
	api       *game_interface.GameInterface
)

func main() {
	tA := time.Now()

	SystemTestingLogin()
	defer func() {
		c.Conn().Close()
		time.Sleep(time.Second)
	}()

	SystemTestingCommands()
	SystemTestingUUIDSafeString()
	SystemTestingSetblock()
	SystemTestingReplaceitem()
	SystemTestingQuerytarget()
	SystemTestingStructrueBackup()
	SystemTestingBotClick()
	SystemTestingContainer()
	SystemTestingItemStackOperation()
	SystemTestingItemCopy()
	SystemTestingItemTransition()

	pterm.Success.Printfln("System Testing: ALL PASS (Time used = %v)", time.Since(tA))
}

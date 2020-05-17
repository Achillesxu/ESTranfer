package core

import (
	"fmt"
	"go.uber.org/zap"
	"os"
)

func Start(app *ApplicationContext, exitChannel chan os.Signal) int {

	// Validate that the ApplicationContext is complete
	if (app == nil) || (app.Logger == nil) || (app.LogLevel == nil) {
		// Didn't get a valid ApplicationContext, so we'll set up our own, with the logger
		app = &ApplicationContext{}
		app.Logger, app.LogLevel = ConfigureLogger()
		defer app.Logger.Sync()
	}
	app.Logger.Info("Started ESTransfer")

	// Set up a specific child logger for main
	log := app.Logger.With(zap.String("type", "main"), zap.String("name", "estransfer"))

	// start main entrance func
	// TODO
	fmt.Println("we ready to deal with data")

	// Wait until we're told to exit
	<-exitChannel
	log.Info("Shutdown triggered")

	// clear all things

	// Exit cleanly
	return 0
}

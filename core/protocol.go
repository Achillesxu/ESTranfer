package core

import (
	"go.uber.org/zap"
)

// ApplicationContext is a structure that holds objects that are used across all coordinators and modules. This is
// used in lieu of passing individual arguments to all functions.
type ApplicationContext struct {
	// Logger is a configured zap.Logger instance. It is to be used by the main routine directly, and the main routine
	// creates loggers for each of the coordinators to use that have fields set to identify that coordinator.
	//
	// This field can be set prior to calling core.Start() in order to pre-configure the logger. If it is not set,
	// core.Start() will set up a default logger using the application config.
	Logger *zap.Logger

	// LogLevel is an AtomicLevel instance that has been used to set the default level of the Logger. It is used to
	// dynamically adjust the logging level (such as via an HTTP call)
	//
	// If Logger has been set prior to calling core.Start(), LogLevel must be set as well.
	LogLevel *zap.AtomicLevel

	// This is used by the main Burrow routines to signal that the configuration is valid. The rest of the code should
	// not care about this, as the application will exit if the configuration is not valid.
	ConfigurationValid bool
}

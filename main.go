package main

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

// exitCode wraps a return value for the application
type exitCode struct{ Code int }

func handleExit() {
	if e := recover(); e != nil {
		if exit, ok := e.(exitCode); ok {
			if exit.Code != 0 {
				_, _ = fmt.Fprintln(os.Stderr, "ESTransfer failed at", time.Now().Format("January 2, 2006 at 3:04pm (MST)"))
			} else {
				_, _ = fmt.Fprintln(os.Stderr, "Stopped ESTransfer at", time.Now().Format("January 2, 2006 at 3:04pm (MST)"))
			}

			os.Exit(exit.Code)
		}
		panic(e) // not an exitCode, bubble up
	}
}

func main() {
	// This makes sure that we panic and run defers correctly
	defer handleExit()

	runtime.GOMAXPROCS(runtime.NumCPU())
	// The only command line arg is the config file
	configPath := flag.String("config-dir", "./config", "Directory that contains the configuration file")
	flag.Parse()

	// Load the configuration from the file
	viper.SetConfigName("estransfer")
	viper.SetConfigType("toml")
	viper.AddConfigPath(*configPath)
	_, _ = fmt.Fprintln(os.Stderr, "Reading configuration from", *configPath)

	err := viper.ReadInConfig()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Failed reading configuration:", err.Error())
		panic(exitCode{1})
	}
	stdoutLogfile := viper.GetString("general.stdout-logfile")

	fmt.Println(stdoutLogfile)

	exitChannel := make(chan os.Signal, 1)
	signal.Notify(exitChannel, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	// This triggers handleExit (after other defers), which will then call os.Exit properly
	panic(exitCode{core.Start(nil, exitChannel)})
}

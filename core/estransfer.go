package core

import (
	"fmt"
	"os"
	"time"
)

func Start(exitChannel chan os.Signal) int {
	time.Sleep(10 * time.Second)
	fmt.Println("hello world")
	// Exit cleanly
	return 0
}

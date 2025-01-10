package helper

import (
	"fmt"
	"runtime"
	"time"
)

func LogLine(message string) {
	pc, _, _, _ := runtime.Caller(1)
	funcName := runtime.FuncForPC(pc).Name()
	currentTime := time.Now().Format("15:04:05")
	fmt.Printf("\033[1;34m[%s]\033[0m/\033[1;32m[%s]\033[0m: %s\n", funcName, currentTime, message)
}

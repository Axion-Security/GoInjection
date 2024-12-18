package helper

import (
	"fmt"
	"github.com/fatih/color" // Colors
	"runtime"
)

func LogError(err error) {
	if err != nil {
		pc, filename, line, _ := runtime.Caller(1)

		funcName := runtime.FuncForPC(pc).Name()
		sourceInfo := fmt.Sprintf("[%s:%d]", filename, line)

		os := runtime.GOOS
		goVersion := runtime.Version()

		fmt.Printf("╭─ [%s] in %s %s %s\n",
			color.New(color.FgRed, color.Bold).Sprint("error"),
			color.New(color.FgBlue).Sprint(funcName),
			color.New(color.FgGreen).Sprint(sourceInfo),
			color.New(color.FgCyan).Sprint("↙"),
		)
		fmt.Printf("│   %v\n", err)
		fmt.Printf("│\n")
		fmt.Printf("│  %s Something went wrong! \n", color.New(color.FgCyan).Sprint("↑"))
		fmt.Printf("│  %s Please check the error message above for details.\n", color.New(color.FgCyan).Sprint("→"))
		fmt.Printf("│  %s User OS: %s\n", color.New(color.FgCyan).Sprint("→"), os)
		fmt.Printf("│  %s Go Version: %s\n", color.New(color.FgCyan).Sprint("→"), goVersion)
		fmt.Printf("│  %s User OS and Go Version information omitted for privacy.\n", color.New(color.FgCyan).Sprint("→"))
		fmt.Println("╰───────────────────────────────────")
	}
}

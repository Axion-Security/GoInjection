package helper

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/fatih/color"
)

func ClearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	LogError(err)
}

func ReadKey() {
	_, _ = fmt.Scanln()
}

func WriteLine(option, value string, newLine, padding bool) {
	gray := color.New(color.FgHiBlack).SprintFunc()
	white := color.New(color.FgHiWhite).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	magenta := color.New(color.FgMagenta).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	if padding {
		fmt.Print("  ")
	}

	fmt.Print(gray("["))

	var optionColor string
	switch option {
	case "!":
		optionColor = cyan(option)
	case "?":
		optionColor = yellow(option)
	case ">":
		optionColor = magenta(option)
	case "+":
		optionColor = green(option)
	case "-":
		optionColor = red(option)
	default:
		optionColor = gray(option)
	}

	fmt.Print(optionColor)
	fmt.Print(gray("] "))

	if option == ">" {
		fmt.Print(white(value))
	} else {
		if newLine {
			fmt.Println(white(value))
		} else {
			fmt.Print(white(value))
		}
	}
}

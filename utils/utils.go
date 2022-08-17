package utils

import (
	"fmt"
	"os"
	"github.com/fatih/color"
)

// Always catch on finish
func OnFinish() {
	fmt.Print("Press [enter] to continue ...")
	fmt.Scanln()
}

// Print error in red and then exit with exit status 1
func ErrorFatal(err error) {
	red := color.New(color.FgRed).SprintFunc()
	fmt.Println(red(err))
	os.Exit(1)
}

// Print log in yellow as warning
func Warn(msg string) {
	yellow := color.New(color.FgYellow).SprintFunc()
	fmt.Println(yellow(msg))
}

// Print success in green
func Green(msg string) {
	green := color.New(color.Bold, color.FgGreen).SprintFunc()
	fmt.Println(green(msg))
}

// Print header in bold font
func Header(msg string) {
	bold := color.New(color.Bold).SprintFunc()
	fmt.Println(bold(msg))
}

// Print Title in bold cyan font
func Title(msg string) {
	boldCyan := color.New(color.Bold, color.FgCyan).SprintFunc()
	fmt.Println(boldCyan(msg))
}
package utils

import (
	"fmt"
	"github.com/fatih/color"
	"golang.org/x/sys/windows"
	"os"
	"runtime"
)

func FixWindowsColors() {
	if runtime.GOOS == "windows" {
		stdout := windows.Handle(os.Stdout.Fd())
		var originalMode uint32

		windows.GetConsoleMode(stdout, &originalMode)
		windows.SetConsoleMode(stdout, originalMode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)
	}
}

// Print error in red and then exit with exit status 1
// Escape colors work differntly in Windows. Swapping to tabs/spaces.
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

// Print in in bold cyan font
func BoldCyan(msg string) {
	boldCyan := color.New(color.Bold, color.FgCyan).SprintFunc()
	fmt.Println(boldCyan(msg))
}

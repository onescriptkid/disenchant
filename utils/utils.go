package utils

import (
	"fmt"
	"os"
	"github.com/fatih/color"
	"runtime"
	"errors"
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

// Print in in bold cyan font
func BoldCyan(msg string) {
	boldCyan := color.New(color.Bold, color.FgCyan).SprintFunc()
	fmt.Println(boldCyan(msg))
}

// Get LockFile Locations
func GetLockFilePaths()(paths []string, err error){
	osname := runtime.GOOS
    switch osname {
	case "windows":
		samePath := "./lockfile"
		lolPath :=  "C:\\Riot Games\\League of Legends\\lockfile"
		pbePath := "C:\\Riot Games\\League of Legends (PBE)\\lockfile"
		paths = []string{ samePath, lolPath, pbePath }
		return
    case "darwin":
		err = errors.New("Unavailable on OSX")
		return
    case "linux":
		err = errors.New("Unavailable on linux")
		return
    default:
		fmt.Printf("%s.\n", osname)
		msg := fmt.Sprintf("Unavailabe on %s", osname)
		err = errors.New(msg)
		return
    }
}
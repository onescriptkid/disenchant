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
	fmt.Scanln()
}

// Print error in red and then exit with exit status 1
// Escape colors work differntly in Windows. Swapping to tabs/spaces.
func ErrorFatal(err error) {
	osname := runtime.GOOS
    switch osname {
	case "windows":
		fmt.Println(err)
    case "darwin":
		red := color.New(color.FgRed).SprintFunc()
		fmt.Println(red(err))
    case "linux":
		red := color.New(color.FgRed).SprintFunc()
		fmt.Println(red(err))
    default:
		fmt.Println(err)
    }
	os.Exit(1)
}

// Print log in yellow as warning
func Warn(msg string) {
	osname := runtime.GOOS
	switch osname {
	case "windows":
		fmt.Println(msg)
    case "darwin":
		yellow := color.New(color.FgYellow).SprintFunc()
		fmt.Println(yellow(msg))
    case "linux":
		yellow := color.New(color.FgYellow).SprintFunc()
		fmt.Println(yellow(msg))
    default:
		fmt.Println(msg)
    }
}

// Print success in green
func Green(msg string) {
	osname := runtime.GOOS
	switch osname {
	case "windows":
		fmt.Println(msg)
    case "darwin":
		green := color.New(color.Bold, color.FgGreen).SprintFunc()
		fmt.Println(green(msg))
    case "linux":
		green := color.New(color.Bold, color.FgGreen).SprintFunc()
		fmt.Println(green(msg))
    default:
		fmt.Println(msg)
    }
}

// Print header in bold font
func Header(msg string) {
	osname := runtime.GOOS
	switch osname {
	case "windows":
		fmt.Println(msg)
    case "darwin":
		bold := color.New(color.Bold).SprintFunc()
		fmt.Println(bold(msg))
    case "linux":
		bold := color.New(color.Bold).SprintFunc()
		fmt.Println(bold(msg))
    default:
		fmt.Println(msg)
    }
}

// Print Title in bold cyan font
func Title(msg string) {
	osname := runtime.GOOS
	switch osname {
	case "windows":
		fmt.Println(msg)
    case "darwin":
		boldCyan := color.New(color.Bold, color.FgCyan).SprintFunc()
		fmt.Println(boldCyan(msg))
    case "linux":
		boldCyan := color.New(color.Bold, color.FgCyan).SprintFunc()
		fmt.Println(boldCyan(msg))
    default:
		fmt.Println(msg)
    }
}

// Print in in bold cyan font
func BoldCyan(msg string) {
	osname := runtime.GOOS
	switch osname {
	case "windows":
		fmt.Println(msg)
    case "darwin":
		boldCyan := color.New(color.Bold, color.FgCyan).SprintFunc()
		fmt.Println(boldCyan(msg))
    case "linux":
		boldCyan := color.New(color.Bold, color.FgCyan).SprintFunc()
		fmt.Println(boldCyan(msg))
    default:
		fmt.Println(msg)
    }
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
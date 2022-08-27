package utils

import (
	"errors"
	"fmt"
	"runtime"
)

// Always catch on finish
func OnFinish() {
	fmt.Print("Press [enter] to continue ...")
	fmt.Scanln()
}

// Get LockFile Locations
func GetLockFilePaths() (paths []string, err error) {
	osname := runtime.GOOS
	switch osname {
	case "windows":
		samePath := "./lockfile"
		lolPath := "C:\\Riot Games\\League of Legends\\lockfile"
		pbePath := "C:\\Riot Games\\League of Legends (PBE)\\lockfile"
		paths = []string{samePath, lolPath, pbePath}
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

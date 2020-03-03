package config

import (
	"os"
	"os/user"
	"path/filepath"
	"runtime"
)

// defaultConfigDir builds the default expected configuration path
func defaultConfigDir() string {
	// Try to place the data folder in the user's home dir
	home := homeDir()

	// If we can not guess a home dir, just fall back to current directory
	if "" == home {
		return "."
	}

	// Set the default config path for different
	switch runtime.GOOS {
	case "darwin":
		return filepath.Join(home, "Library", "FantomApi")
	case "windows":
		return filepath.Join(home, "FantomApi")
	default:
		return filepath.Join(home, ".fantomapi")
	}
}

// homeDir guess current user's home dir
func homeDir() string {
	// Unix like env variable
	if home := os.Getenv("HOME"); home != "" {
		return home
	}

	// Windows like env variable; available only in User context
	if home := os.Getenv("USERPROFILE"); home != "" {
		return home
	}

	if usr, err := user.Current(); err == nil {
		return usr.HomeDir
	}

	return ""
}

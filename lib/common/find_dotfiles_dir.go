package common

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/noclaps/dot/lib/common/log"
	. "github.com/noclaps/dot/lib/types"
)

func FindDotfilesDir() AbsolutePath {
	dotfilesDir, err := findDotfilesDir()
	if err != nil {
		log.Fatal("Error finding dotfiles directory: %v", err)
	}
	if !filepath.IsAbs(dotfilesDir) {
		log.Fatal("Dotfiles directory must be an absolute path: %s", dotfilesDir)
	}
	log.Info("Using dotfiles directory: %s", dotfilesDir)
	return NewAbsolutePath(dotfilesDir)
}

func findDotfilesDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error retrieving home directory: %v", err)
	}

	// Try ~/.dotfiles
	dotfilesDir := filepath.Join(homeDir, ".dotfiles")
	if fileInfo, err := os.Stat(dotfilesDir); err == nil && fileInfo.IsDir() {
		return dotfilesDir, nil
	}

	err = fmt.Errorf("Dotfiles directory does not exist: %s", filepath.Join(homeDir, ".dotfiles"))
	return "", err
}

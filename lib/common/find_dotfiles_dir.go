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
	// 1. Try $DOT_DIR if defined
	if dotDir := os.Getenv(ENV_DOT_DIR); dotDir != "" {
		fileInfo, err := os.Stat(dotDir)
		if err == nil && fileInfo.IsDir() {
			return dotDir, nil
		}
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error retrieving home directory: %v", err)
	}

	// 2. Try $XDG_DATA_HOME/dotfiles (or ~/.local/share/dotfiles)
	xdgDataHome := os.Getenv(ENV_XDG_DATA_HOME)
	if xdgDataHome == "" {
		xdgDataHome = filepath.Join(homeDir, ".local", "share")
	}
	dotfilesDir := filepath.Join(xdgDataHome, "dotfiles")
	if fileInfo, err := os.Stat(dotfilesDir); err == nil && fileInfo.IsDir() {
		return dotfilesDir, nil
	}

	// 3. Try ~/.dotfiles
	dotfilesDir = filepath.Join(homeDir, ".dotfiles")
	if fileInfo, err := os.Stat(dotfilesDir); err == nil && fileInfo.IsDir() {
		return dotfilesDir, nil
	}

	err = fmt.Errorf("none of the candidate dotfiles directories exist:\n  - $DOT_DIR = '%s'\n  - %s\n  - %s",
		os.Getenv(ENV_DOT_DIR),
		filepath.Join(xdgDataHome, "dotfiles"),
		filepath.Join(homeDir, ".dotfiles"))
	return "", err
}

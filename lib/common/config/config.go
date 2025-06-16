package config

import (
	"os"

	"github.com/noclaps/dot/lib/common/log"
)

type Config struct {
	// Where to install the symlinks. In most cases this will be either "$HOME"
	// (dotfiles) or "/" (root configs). Must be an absolute path. It can contain
	// environment variables.
	TargetDir string
	// Files and directories to ignore. Each entry is a glob pattern relative to
	// the dotfiles directory. IMPORTANT: Hidden files/directories are ignored by
	// default. If you set `implicit_dot` to false, you should remove the `**/.*`
	// pattern from this list.
	ExcludeFiles []string
	// Files and directories that are always symlinked, overriding
	// `exclude_files`. Each entry is a glob pattern relative to the dotfiles
	// directory.
	IncludeFiles []string
	// You can get a large performance boost by setting this to `false`, but read
	// this first: https://github.com/pol-rivero/doot/wiki/Tip:-set-explore_excluded_dirs-to-false
	ExploreExcludedDirs bool
	// If set to true, files and directories in the root of the dotfiles directory
	// will be prefixed with a dot. For example, `<dotfiles dir>/config/foo` will
	// be symlinked to `~/.config/foo`. This is useful if you don't want to have
	// hidden files in the root of the dotfiles directory.
	ImplicitDot bool
	// Top-level files and directories that won't be prefixed with a dot if
	// `implicit_dot` is set to true. Each entry is the name of a file or
	// directory in the root of the dotfiles directory.
	ImplicitDotIgnore []string
}

func GetConfig() Config {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Error retrieving home directory: %v", err)
	}
	return Config{
		TargetDir:           homedir,
		ExcludeFiles:        []string{"**/.*", "LICENSE", "README.md"},
		IncludeFiles:        []string{},
		ExploreExcludedDirs: false,
		ImplicitDot:         true,
		ImplicitDotIgnore:   []string{},
	}
}

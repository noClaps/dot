package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/noclaps/dot/lib/common/log"
	. "github.com/noclaps/dot/lib/types"
	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	TargetDir           string   `toml:"target_dir"`
	ExcludeFiles        []string `toml:"exclude_files"`
	IncludeFiles        []string `toml:"include_files"`
	ExploreExcludedDirs bool     `toml:"explore_excluded_dirs"`
	ImplicitDot         bool     `toml:"implicit_dot"`
	ImplicitDotIgnore   []string `toml:"implicit_dot_ignore"`
}

func DefaultConfig() Config {
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

func FromFile(path AbsolutePath) Config {
	config := DefaultConfig()
	fileContents, err := os.ReadFile(path.Str())
	if err != nil {
		log.Info("Config file not found or unaccessible, using default config")
		return config
	}
	err = toml.Unmarshal(fileContents, &config)
	if err != nil {
		log.Error("Error parsing config file: %v", err)
	}
	verifyConfig(&config)
	return config
}

func FromDotfilesDir(dotfilesDir AbsolutePath) Config {
	return FromFile(dotfilesDir.Join("dot").Join("config.toml"))
}

func verifyConfig(config *Config) {
	config.TargetDir = filepath.Clean(os.ExpandEnv(config.TargetDir))
	if !filepath.IsAbs(config.TargetDir) {
		log.Fatal("Invalid config: 'target_dir = %s', must be an absolute path", config.TargetDir)
	}
	for _, implicitDotIgnore := range config.ImplicitDotIgnore {
		if strings.ContainsRune(implicitDotIgnore, filepath.Separator) {
			topLevelDir := RelativePath(implicitDotIgnore).TopLevelDir()
			log.Fatal("Invalid config. 'implicit_dot_ignore -> %s' must be a top-level file or directory. Consider adding '%s' instead",
				implicitDotIgnore, topLevelDir)
		}
	}
}

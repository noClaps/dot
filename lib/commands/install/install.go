package install

import (
	"path/filepath"

	"github.com/noclaps/dot/lib/common"
	"github.com/noclaps/dot/lib/common/cache"
	"github.com/noclaps/dot/lib/common/config"
	"github.com/noclaps/dot/lib/linkmode"
	. "github.com/noclaps/dot/lib/types"
)

type GetFilesFunc func(*config.Config, AbsolutePath) []RelativePath

func Install() {
	getFiles := func(config *config.Config, dotfilesDir AbsolutePath) []RelativePath {
		filter := CreateFilter(config)
		return ScanDirectory(dotfilesDir, &filter)
	}
	installImpl(getFiles)
}

func Clean() {
	getFiles := func(config *config.Config, dotfilesDir AbsolutePath) []RelativePath {
		return []RelativePath{}
	}
	installImpl(getFiles)
}

func installImpl(getFiles GetFilesFunc) {
	dotfilesDir := common.FindDotfilesDir()
	config := config.GetConfig()
	linkMode := linkmode.GetLinkMode(&config)

	cacheKey := dotfilesDir.Str() + string(filepath.ListSeparator) + config.TargetDir
	cache := cache.Load()
	installedFilesCache := cache.GetEntry(cacheKey)
	if fullClean {
		installedFilesCache.Links = linkMode.RecalculateCache(dotfilesDir, config.TargetDir)
	}

	fileList := getFiles(&config, dotfilesDir)
	fileMapping := NewFileMapping(dotfilesDir, &config, fileList)

	oldLinks := installedFilesCache.GetLinks()
	added := fileMapping.InstallNewLinks()
	removed := fileMapping.RemoveStaleLinks(&oldLinks)

	installedFilesCache.SetLinks(fileMapping.GetInstalledTargets())
	cache.Save()

	printChanges(added, removed)
}

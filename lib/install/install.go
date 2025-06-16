package install

import (
	"path/filepath"

	"github.com/noclaps/dot/lib/common"
	"github.com/noclaps/dot/lib/common/cache"
	"github.com/noclaps/dot/lib/common/config"
	. "github.com/noclaps/dot/lib/types"
)

type GetFilesFunc func(*config.Config, AbsolutePath) []RelativePath

func Install(fullClean bool) {
	getFiles := func(config *config.Config, dotfilesDir AbsolutePath) []RelativePath {
		filter := CreateFilter(config)
		return ScanDirectory(dotfilesDir, &filter)
	}
	installImpl(getFiles, fullClean)
}

func Clean(fullClean bool) {
	getFiles := func(config *config.Config, dotfilesDir AbsolutePath) []RelativePath {
		return []RelativePath{}
	}
	installImpl(getFiles, fullClean)
}

func installImpl(getFiles GetFilesFunc, fullClean bool) {
	dotfilesDir := common.FindDotfilesDir()
	config := config.FromDotfilesDir(dotfilesDir)

	cacheKey := dotfilesDir.Str() + string(filepath.ListSeparator) + config.TargetDir
	cache := cache.Load()
	installedFilesCache := cache.GetEntry(cacheKey)
	if fullClean {
		recalculateCache(installedFilesCache, dotfilesDir, config.TargetDir)
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

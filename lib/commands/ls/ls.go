package ls

import (
	"github.com/noclaps/dot/lib/common"
	"github.com/noclaps/dot/lib/common/cache"
	"github.com/noclaps/dot/lib/common/config"
	"github.com/noclaps/dot/lib/common/log"
)

func ListInstalledFiles() {
	dotfilesDir := common.FindDotfilesDir()
	config := config.GetConfig()

	cacheKey := cache.ComputeCacheKey(dotfilesDir, config.TargetDir)
	cache := cache.Load()
	installedFilesCache := cache.GetEntry(cacheKey)

	installedLinks := installedFilesCache.GetLinks()
	log.Printlnf("%s", installedLinks.PrintList())
}

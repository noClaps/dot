package ls

import (
	"path/filepath"

	"github.com/noclaps/dot/lib/common"
	"github.com/noclaps/dot/lib/common/cache"
	"github.com/noclaps/dot/lib/common/config"
	"github.com/noclaps/dot/lib/common/log"
)

func ListInstalledFiles(asJson bool) {
	dotfilesDir := common.FindDotfilesDir()
	config := config.GetConfig()

	cacheKey := dotfilesDir.Str() + string(filepath.ListSeparator) + config.TargetDir
	cache := cache.Load()
	installedFilesCache := cache.GetEntry(cacheKey)

	installedLinks := installedFilesCache.GetLinks()
	if asJson {
		log.Printlnf("%s", installedLinks.ToJson())
	} else {
		log.Printlnf("%s", installedLinks.PrintList())
	}
}

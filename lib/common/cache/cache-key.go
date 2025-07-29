package cache

import (
	"path/filepath"

	. "github.com/noclaps/dot/lib/types"
)

func ComputeCacheKey(dotfilesDir AbsolutePath, targetDir string) string {
	return dotfilesDir.Str() + string(filepath.ListSeparator) + targetDir
}

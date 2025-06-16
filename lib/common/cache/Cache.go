package cache

import (
	"os"
	"path"

	"github.com/noclaps/dot/lib/common"
	"github.com/noclaps/dot/lib/common/log"
	. "github.com/noclaps/dot/lib/types"
)

const CURRENT_CACHE_VERSION uint32 = 2

func Load() DotCache {
	fileContents, err := os.ReadFile(getCachePath())
	if err != nil {
		log.Info("Cache read error: %v, creating new cache", err)
		return DotCache{
			Version: CURRENT_CACHE_VERSION,
			Entries: []*CacheEntry{},
		}
	}

	var cacheData DotCache
	err = cacheData.UnmarshalBinary(fileContents)
	if err != nil {
		log.Warning("Error parsing cache file: %v, creating new cache", err)
		return DotCache{
			Version: CURRENT_CACHE_VERSION,
			Entries: []*CacheEntry{},
		}
	}

	if cacheData.Version != CURRENT_CACHE_VERSION {
		log.Info("Cache version mismatch: expected %d, got %d", CURRENT_CACHE_VERSION, cacheData.Version)
		return DotCache{
			Version: CURRENT_CACHE_VERSION,
			Entries: []*CacheEntry{},
		}
	}

	return cacheData
}

func (cache *DotCache) Save() {
	marshalledData, err := cache.MarshalBinary()
	if err != nil {
		log.Error("Error marshalling cache data: %v", err)
		return
	}

	err = os.WriteFile(getCachePath(), marshalledData, 0644)
	if err != nil {
		log.Error("Error saving cache file: %v", err)
	}
}

func (cache *DotCache) GetEntry(cacheKey string) *InstalledFilesCache {
	for _, entry := range cache.Entries {
		if entry.CacheKey == cacheKey {
			return entry.InstalledFiles
		}
	}

	newEntry := CacheEntry{
		CacheKey:       cacheKey,
		InstalledFiles: &InstalledFilesCache{},
	}
	cache.Entries = append(cache.Entries, &newEntry)
	return newEntry.InstalledFiles
}

func (filesCache *InstalledFilesCache) GetLinks() SymlinkCollection {
	links := NewSymlinkCollection(len(filesCache.Links))
	for _, link := range filesCache.Links {
		links.Add(NewAbsolutePath(link.Path), NewAbsolutePath(link.Content))
	}
	return links
}

func (filesCache *InstalledFilesCache) SetLinks(links SymlinkCollection) {
	filesCache.Links = make([]*InstalledFile, 0, links.Len())
	for path, content := range links.Iter() {
		filesCache.Links = append(filesCache.Links, &InstalledFile{
			Path:    path.Str(),
			Content: content.Str(),
		})
	}
}

func getCachePath() string {
	cacheDir := getCacheContainingDir()
	err := os.MkdirAll(cacheDir, 0755)
	if err != nil {
		log.Fatal("Error creating cache directory: %v", err)
	}
	return path.Join(cacheDir, "dot-cache.bin")
}

func getCacheContainingDir() string {
	cacheDir := os.Getenv(common.ENV_DOT_CACHE_DIR)
	if cacheDir != "" {
		return cacheDir
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Error retrieving home directory: %v", err)
	}
	return path.Join(homeDir, ".cache", "dot")
}

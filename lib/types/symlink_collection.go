package types

import (
	"slices"
	"strings"
)

type SymlinkCollection struct {
	// link path -> link content (target)
	links map[AbsolutePath]AbsolutePath
}

func NewSymlinkCollection(capacity int) SymlinkCollection {
	return SymlinkCollection{make(map[AbsolutePath]AbsolutePath, capacity)}
}

func (sc *SymlinkCollection) Add(linkPath, linkContent AbsolutePath) {
	sc.links[linkPath] = linkContent
}

func (sc *SymlinkCollection) Len() int {
	return len(sc.links)
}

func (sc *SymlinkCollection) Iter() map[AbsolutePath]AbsolutePath {
	return sc.links
}

func (sc *SymlinkCollection) PrintList() string {
	paths := make([]string, 0, len(sc.links))
	for path := range sc.links {
		paths = append(paths, path.Str())
	}
	slices.Sort(paths)

	var sb strings.Builder
	for _, pathStr := range paths {
		path := AbsolutePath(pathStr)
		content := sc.links[path]
		sb.WriteString(path.Str())
		sb.WriteString(" -> ")
		sb.WriteString(content.Str())
		sb.WriteString("\n")
	}
	return sb.String()
}

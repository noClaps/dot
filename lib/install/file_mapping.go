package install

import (
	"os"
	"slices"
	"strings"

	"github.com/noclaps/dot/lib/common/config"
	"github.com/noclaps/dot/lib/common/log"
	. "github.com/noclaps/dot/lib/types"
	"github.com/noclaps/dot/lib/utils"
	"github.com/noclaps/dot/lib/utils/files"
	"github.com/noclaps/dot/lib/utils/optional"
	"github.com/noclaps/dot/lib/utils/set"
)

type SourcePath struct {
	path AbsolutePath
}

type FileMapping struct {
	mapping           map[AbsolutePath]SourcePath // Installed target (symlink path) -> source (dotfile, symlink content/target)
	sourceBaseDir     AbsolutePath
	targetBaseDir     AbsolutePath
	implicitDot       bool
	implicitDotIgnore set.Set[string]
	targetsSkipped    []AbsolutePath
}

func NewFileMapping(dotfilesDir AbsolutePath, config *config.Config, sourceFiles []RelativePath) FileMapping {
	mapping := FileMapping{
		mapping:           make(map[AbsolutePath]SourcePath, len(sourceFiles)),
		sourceBaseDir:     dotfilesDir,
		targetBaseDir:     NewAbsolutePath(config.TargetDir),
		implicitDot:       config.ImplicitDot,
		implicitDotIgnore: set.NewFromSlice(config.ImplicitDotIgnore),
		targetsSkipped:    make([]AbsolutePath, 0),
	}
	for _, sourceFile := range sourceFiles {
		mapping.Add(sourceFile)
	}
	return mapping
}

func (fm *FileMapping) Add(relativeSource RelativePath) {
	relativeTarget := fm.mapSourceToTarget(relativeSource)
	if relativeTarget.IsEmpty() {
		return
	}
	source := fm.sourceBaseDir.JoinPath(relativeSource)
	target := fm.targetBaseDir.JoinPath(relativeTarget.Value())

	oldSource, oldSourceExists := fm.mapping[target]
	preferNewSource := !oldSourceExists
	if preferNewSource {
		fm.mapping[target] = SourcePath{
			path: source,
		}
	} else {
		// This is rare, but it can happen if 2 files map to the same target after adding the implicit dot
		log.Warning("Conflicting files: %s and %s both map to %s. Ignoring %s", oldSource.path, source, target, source)
	}
}

func (fm *FileMapping) GetInstalledTargets() SymlinkCollection {
	targets := NewSymlinkCollection(len(fm.mapping))
	for targetPath, sourcePath := range fm.mapping {
		if !slices.Contains(fm.targetsSkipped, targetPath) {
			targets.Add(targetPath, sourcePath.path)
		}
	}
	return targets
}

func (fm *FileMapping) InstallNewLinks() []AbsolutePath {
	createdLinks := make([]AbsolutePath, 0, 5)
	for target, sourceStruct := range fm.mapping {
		newSource := sourceStruct.path
		if earlySkipTarget(target, newSource) {
			continue
		}

		fileInfo, err := os.Lstat(target.Str())
		if err == nil {
			added := fm.handleTargetAlreadyExists(fileInfo, target, newSource)
			if added {
				createdLinks = append(createdLinks, target)
			}
			continue
		}
		if os.IsNotExist(err) && files.EnsureParentDir(target) {
			log.Info("Linking %s -> %s", target, newSource)
			err = os.Symlink(newSource.Str(), target.Str())
			if err == nil {
				createdLinks = append(createdLinks, target)
				continue
			}
		}
		log.Error("Failed to create link %s -> %s: %s", target, newSource, err)
	}
	return createdLinks
}

func earlySkipTarget(target AbsolutePath, newSource AbsolutePath) bool {
	linkSource, linkErr := os.Readlink(target.Str())
	if linkErr != nil {
		// Either the link doesn't exist or it's not a symlink. In either case, proceed with the installation
		return false
	}
	if linkSource == newSource.Str() {
		log.Info("Link %s is already correct", target)
		return true
	}
	return false
}

func (fm *FileMapping) RemoveStaleLinks(previousLinks *SymlinkCollection) []AbsolutePath {
	removedLinks := make([]AbsolutePath, 0, 5)
	for previousLinkPath := range previousLinks.Iter() {
		if _, contains := fm.mapping[previousLinkPath]; !contains {
			if !canBeSafelyRemoved(previousLinkPath, fm.sourceBaseDir) {
				log.Info("%s appears to have been modified externally. Skipping removal to avoid data loss.", previousLinkPath)
				continue
			}
			log.Info("Removing link %s", previousLinkPath)
			success := files.RemoveAndCleanup(previousLinkPath, fm.targetBaseDir)
			if success {
				removedLinks = append(removedLinks, previousLinkPath)
			}
		}
	}
	return removedLinks
}

func (fm *FileMapping) handleTargetAlreadyExists(fileInfo os.FileInfo, target, source AbsolutePath) bool {
	if fileInfo.Mode()&os.ModeSymlink != 0 {
		return fm.handleExistingSymlink(target, source)
	} else if fileInfo.Mode().IsRegular() {
		return fm.handleExistingFile(target, source)
	} else {
		log.Warning("Target %s exists but is not a symlink or a regular file, skipping", target)
		return false
	}
}

func (fm *FileMapping) handleExistingSymlink(target, source AbsolutePath) bool {
	linkSource, linkErr := os.Readlink(target.Str())
	if linkErr != nil {
		log.Error("Failed to read link %s: %s", target, linkErr)
		return false
	}
	if strings.HasPrefix(linkSource, fm.sourceBaseDir.Str()) {
		log.Info("Link %s is incorrect (%s) but points to the source directory, replacing silently with %s", target, linkSource, source)
		err := files.ReplaceWithSymlink(target, source)
		return err == nil
	}
	replace := utils.RequestInput("yN", "Link %s already exists, but it points to %s instead of %s. Replace it?", target, linkSource, source)
	if replace == 'y' {
		err := files.ReplaceWithSymlink(target, source)
		return err == nil
	} else {
		fm.targetsSkipped = append(fm.targetsSkipped, target)
		return false
	}
}

func (fm *FileMapping) handleExistingFile(target, source AbsolutePath) bool {
	contents, readErr := os.ReadFile(target.Str())
	if readErr != nil {
		log.Error("Failed to read file %s: %s", target, readErr)
		return false
	}
	sourceContents, readErr := os.ReadFile(source.Str())
	if readErr != nil {
		log.Error("Failed to read file %s: %s", source, readErr)
		return false
	}
	if string(contents) == string(sourceContents) {
		log.Info("File %s exists but its contents are identical to %s, replacing silently", target, source)
		err := files.ReplaceWithSymlink(target, source)
		return err == nil
	}
	for {
		replace := utils.RequestInput("yNa", "File %s already exists, but its contents differ from %s. Replace it? (A to adopt changes into dotfiles repo)", target, source)
		switch replace {
		case 'y':
			err := files.ReplaceWithSymlink(target, source)
			return err == nil
		case 'n':
			fm.targetsSkipped = append(fm.targetsSkipped, target)
			return false
		case 'a':
			err := files.AdoptChanges(target, source)
			return err == nil
		}
	}
}

func (fm *FileMapping) mapSourceToTarget(source RelativePath) optional.Optional[RelativePath] {
	target := source
	if fm.implicitDot && !fm.implicitDotIgnore.Contains(source.TopLevelDir()) && !strings.HasPrefix(target.Str(), ".") {
		target = "." + target
	}
	return optional.WrapString(target)
}

func canBeSafelyRemoved(linkPath AbsolutePath, expectedDestinationDir AbsolutePath) bool {
	linkSource, linkErr := os.Readlink(linkPath.Str())
	if linkErr != nil {
		return false
	}
	return strings.HasPrefix(linkSource, expectedDestinationDir.Str())
}

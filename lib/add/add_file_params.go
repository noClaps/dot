package add

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pol-rivero/doot/lib/common/glob_collection"
	. "github.com/pol-rivero/doot/lib/types"
	"github.com/pol-rivero/doot/lib/utils/set"
)

type ProcessAddedFileParams struct {
	dotfilesDir       string
	targetDir         string
	implicitDot       bool
	implicitDotIgnore set.Set[string]
	includeFiles      glob_collection.GlobCollection
	excludeFiles      glob_collection.GlobCollection
}

func ProcessAddedFile(input string, params ProcessAddedFileParams) (RelativePath, error) {
	fileInfo, err := os.Stat(input)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("this file does not exist")
		}
		return "", err
	}
	if fileInfo.IsDir() {
		return "", fmt.Errorf("it's a directory. Consider adding %s/**/* instead", input)
	}
	cleanAbsFile, err := filepath.Abs(input)
	if err != nil {
		return "", fmt.Errorf("error getting absolute path: %v", err)
	}
	if !strings.HasPrefix(cleanAbsFile, params.targetDir) {
		return "", fmt.Errorf("it's not inside target directory %s", params.targetDir)
	}

	relPath, err := constructRelativePath(cleanAbsFile, params)
	if err != nil {
		return "", fmt.Errorf("error getting relative path: %v", err)
	}

	if params.implicitDot && !params.implicitDotIgnore.Contains(relPath.TopLevelDir()) {
		if relPath.IsHidden() {
			relPath = relPath.Unhide()
		} else {
			return "", fmt.Errorf("its current filename is impossible because implicit_dot is true. Add '%s' to implicit_dot_ignore to fix this", relPath.TopLevelDir())
		}
	}

	if err := checkIsIncluded(relPath, params.includeFiles, params.excludeFiles); err != nil {
		return "", err
	}

	return relPath, nil
}

func constructRelativePath(absPath string, params ProcessAddedFileParams) (RelativePath, error) {
	relPathStr, err := filepath.Rel(params.targetDir, absPath)
	if err != nil {
		return "", err
	}
	parts := strings.Split(relPathStr, string(filepath.Separator))
	if len(parts) == 0 {
		return "", fmt.Errorf("empty relative path")
	}

	return RelativePath(filepath.Join(parts...)), nil
}

func checkIsIncluded(relPath RelativePath, includeFiles glob_collection.GlobCollection, excludeFiles glob_collection.GlobCollection) error {
	if relPath.Str() == "." {
		return nil
	}
	if excludeFiles.Matches(relPath) && !includeFiles.Matches(relPath) {
		return fmt.Errorf("%s matches exclude_files but is not included in include_files", relPath)
	}
	return checkIsIncluded(relPath.Parent(), includeFiles, excludeFiles)
}

package linkmode

import (
	"github.com/noclaps/dot/lib/common/config"
	symlink "github.com/noclaps/dot/lib/linkmode/symlink"
	. "github.com/noclaps/dot/lib/types"
)

type LinkMode interface {
	CreateLink(dotfilesSource, target AbsolutePath) error
	IsInstalledLinkOf(maybeInstalledLinkPath string, dotfilePath AbsolutePath) bool
	CanBeSafelyRemoved(linkPath AbsolutePath, expectedDestinationDir string) bool
}

func GetLinkMode(config *config.Config) LinkMode {
	return &symlink.SymlinkLinkMode{}
}

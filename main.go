package main

import (
	"github.com/noclaps/applause"
	"github.com/noclaps/dot/lib/commands/install"
	"github.com/noclaps/dot/lib/commands/ls"
	"github.com/noclaps/dot/lib/common/log"
)

type Args struct {
	Clean bool `type:"option" short:"C" help:"Remove all symlinks created by dot."`
	List  bool `type:"option" short:"l" help:"List the installed (symlinked) dotfiles."`
}

func main() {
	args := Args{}
	err := applause.Parse(&args)
	if err != nil {
		log.Error("%v", err)
	}
	if args.Clean {
		install.Clean()
		return
	}
	if args.List {
		ls.ListInstalledFiles()
		return
	}

	install.Install()
}

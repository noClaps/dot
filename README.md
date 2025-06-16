# dot

This is a fork of [`doot`](https://github.com/pol-rivero/doot). Chances are, you'll want to use `doot` and not this fork, since this fork is highly opinionated and customised just for my personal use.

A fast, simple and intuitive dotfiles manager that just gets the job done. **[Should you try it?](https://github.com/pol-rivero/doot/wiki/Should-I-use-doot%3F)**

## Install

You can install it from Homebrew:

```sh
brew install noclaps/tap/dot
```

or build it from source using Go:

```sh
go install github.com/noclaps/dot@latest
```

or download one of the prebuilt binaries in [Releases](https://github.com/noClaps/dot/releases).

## Usage

```
USAGE: dot [--clean] [--list]

OPTIONS:
  -C, --clean        Remove all symlinks created by dot.
  -l, --list         List the installed (symlinked) dotfiles.
  -h, --help         Display this help and exit.
```

You can simply run `dot` from anywhere in your system. It will symlink all files in your dotfiles directory to your home directory, creating directories as needed. The subsequent runs will incrementally update the symlinks, adding the new files and directories, and removing references to files that are no longer in the dotfiles directory.

```sh
git clone https://your-dotfiles.git ~/.dotfiles
dot
```

To remove the symlinks, you can use `-C` or `--clean`:

```sh
dot -C
dot --clean
```

You can list the installed dotfiles with `-l` or `--list`:

```sh
dot -l
dot --list
```

You can view the help by using `-h` or `--help`:

```sh
dot -h
dot --help
```

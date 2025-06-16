# dot

This is a fork of [`doot`](https://github.com/pol-rivero/doot). Chances are, you'll want to use `doot` and not this fork, since this fork is highly opinionated and customised just for my personal use.

A fast, simple and intuitive dotfiles manager that just gets the job done. **[Should you try it?](https://github.com/pol-rivero/doot/wiki/Should-I-use-doot%3F)**

## Install

Install `dot` from Homebrew:

```sh
brew install noclaps/tap/dot
```

Make sure to run `brew update && brew upgrade` periodically to keep `dot` up to date.

## Usage

If this is your first time setting up a dotfiles repository, read the [Getting Started](https://github.com/pol-rivero/doot/wiki/Getting-Started) guide.

### Create or update symlinks

Simply run `dot` (or `dot install`) from anywhere in your system. It will symlink all files in your dotfiles directory to your home directory, creating directories as needed.
The subsequent runs will incrementally update the symlinks, adding the new files and directories, and removing references to files that are no longer in the dotfiles directory.

```sh
git clone https://your-dotfiles.git ~/.dotfiles

dot  # Installs or updates the symlinks
```

To remove the symlinks, run:

```sh
dot clean
```

## Dotfiles directory location

`dot` searches for your dotfiles in `$HOME/.dotfiles`.

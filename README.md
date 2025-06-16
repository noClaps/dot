# dot

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

Pass `--full-clean` to the `install` or `clean` commands to search for all symlinks that point to the dotfiles directory, even if they were created by another program. This is useful if you created symlinks manually or your dotfiles installation has somehow become corrupted.

### Add a new file to the dotfiles directory

You could manually move the file to the dotfiles directory and run `dot` to symlink it, but there's a command to do it in one step:

```sh
dot add ./some/file [/other/file ...]
```

- You can undo this operation by running `dot restore <file1> ...`, which will replace the symlink with the original regular file, removing it from the dotfiles repository.

## Dotfiles directory location

`dot` searches for your dotfiles in `$HOME/.dotfiles`.

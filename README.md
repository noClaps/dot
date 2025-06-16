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
git clone https://your-dotfiles.git ~/.dotfiles # (or any other directory)

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

By default, `dot` searches for your dotfiles in commonly used directories. In order of priority, it looks for the first directory that exists:

1. `$DOT_DIR`

2. `$XDG_DATA_HOME/dotfiles` (or `$HOME/.local/share/dotfiles` if `XDG_DATA_HOME` is not set)

3. `$HOME/.dotfiles`

Notice how you can set the `DOT_DIR` environment variable to use any custom directory. The first time you run `dot`, if that variable is not yet defined globally, you can set it inline:

```sh
DOT_DIR=/path/to/your/dotfiles dot
```

After that, if you have set `DOT_DIR` in your shell configuration file (`~/.bashrc` or equivalent), you can just run `dot` as usual.

## Configuration file

`dot` reads an optional configuration file: `<dotfiles dir>/dot/config.toml`. This file won't be symlinked when installing. These are the available options and their default values:

```toml
# Where to install the symlinks. In most cases this will be either "$HOME" (dotfiles) or "/" (root configs).
# Must be an absolute path. It can contain environment variables.
target_dir = "$HOME"

# Files and directories to ignore. Each entry is a glob pattern relative to the dotfiles directory.
# IMPORTANT: Hidden files/directories are ignored by default. If you set `implicit_dot` to false, you should remove the `**/.*` pattern from this list.
exclude_files = [
  "**/.*",
  "LICENSE",
  "README.md",
]

# Files and directories that are always symlinked, overriding `exclude_files`. Each entry is a glob pattern relative to the dotfiles directory.
include_files = []

# You can get a large performance boost by setting this to `false`, but read this first:
# https://github.com/pol-rivero/doot/wiki/Tip:-set-explore_excluded_dirs-to-false
explore_excluded_dirs = true

# If set to true, files and directories in the root of the dotfiles directory will be prefixed with a dot. For example, `<dotfiles dir>/config/foo` will be symlinked to `~/.config/foo`.
# This is useful if you don't want to have hidden files in the root of the dotfiles directory.
implicit_dot = true

# Top-level files and directories that won't be prefixed with a dot if `implicit_dot` is set to true. Each entry is the name of a file or directory in the root of the dotfiles directory.
implicit_dot_ignore = [
  "bin"
]
```

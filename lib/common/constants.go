package common

import "path/filepath"

const DOOT_BACKUP_EXT string = ".doot-backup"
const CUSTOM_COMMANDS_DIR string = "doot" + string(filepath.Separator) + "commands"

const IGNORE_HIDDEN_FILES_GLOB string = "**/.*"

const ENV_DOOT_DIR string = "DOOT_DIR"
const ENV_DOOT_CACHE_DIR string = "DOOT_CACHE_DIR"
const ENV_XDG_DATA_HOME string = "XDG_DATA_HOME"

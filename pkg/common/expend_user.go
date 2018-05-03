package common

import (
	"os/user"
	"path/filepath"
)

func ExpendHomePath(path string) string {
	usr, err := user.Current()
	if err == nil {
		if path[:2] == "~/" {
			path = filepath.Join(usr.HomeDir, path[2:])
		}
	}

	return path
}

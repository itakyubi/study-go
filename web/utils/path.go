package utils

import "os"

func FileExists(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return !fi.IsDir()
}

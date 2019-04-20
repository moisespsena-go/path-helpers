// +build dev

package path_helpers

import (
	"path/filepath"
)

func GetCalledFile(abs ...bool) string {
	pth := GetCalledFileAbs(2)
	if len(abs) == 0 || abs[0] == false {
		return StripGoPath(pth)
	}
	return pth
}

func GetCalledDir(abs ...bool) string {
	dir := filepath.Dir(GetCalledFileAbs(2))
	if len(abs) == 0 || abs[0] == false {
		return StripGoPath(dir)
	}
	return dir
}

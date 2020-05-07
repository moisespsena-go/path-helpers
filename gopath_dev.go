// +build dev

package path_helpers

import (
	"path/filepath"
)

func GetCalledFileUp(up int, abs ...bool) string {
	pth := GetCalledFileAbs(up + 2)
	if len(abs) == 0 || abs[0] == false {
		return StripGoPath(pth)
	}
	return pth
}

func GetCalledDirUp(skip int, abs ...bool) string {
	return filepath.Dir(GetCalledFileUp(skip+1, abs...))
}

func GetCalledFile(abs ...bool) string {
	return GetCalledFileUp(1, abs...)
}

func GetCalledDir(abs ...bool) string {
	return GetCalledDirUp(1, abs...)
}

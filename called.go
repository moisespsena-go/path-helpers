package path_helpers

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func GetCalledFileNameStripGoPath(skip int, abs ...bool) (pth string) {
	defer func() {
		pth = strings.TrimPrefix(pth, string(os.PathSeparator))
	}()
	_, filename, _, ok := runtime.Caller(skip)
	if !ok {
		panic(errors.New("Information unavailable."))
	}
	for _, gp := range GOPATHS {
		if gp.HasSrcDir() {
			if p2 := strings.TrimPrefix(filename, filepath.Join(gp.pth, "src")); len(p2) < len(filename) {
				return p2
			}
		} else {
			if p2 := strings.TrimPrefix(filename, gp.pth); len(p2) < len(filename) {
				return p2
			}
		}
	}
	return filename
}

func GetCalledFile(abs ...bool) string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic(errors.New("runtime.Caller: information unavailable."))
	}
	if len(abs) == 0 || abs[0] == false {
		return StripGoPath(filename)
	}
	return filename
}

func GetCalledDir(abs ...bool) string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic(errors.New("runtime.Caller: information unavailable."))
	}
	filename = filepath.Dir(filename)
	if len(abs) == 0 || abs[0] == false {
		return StripGoPath(filename)
	}
	return filename
}

func GetCalledDirOrError(abs ...bool) string {
	file := GetCalledFileNameStripGoPath(2, abs...)
	if file == "" {
		panic("Invalid dir.")
	}
	return filepath.Dir(file)
}

// +build dev

package path_helpers

import (
	"go/build"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	var (
		pth string
		ok  bool
	)

	paths := make(map[string]interface{})

	if _, err := os.Stat("vendor"); err == nil {
		if abs, err := filepath.Abs("vendor"); err == nil {
			paths[abs] = nil
			GOPATHS = append(GOPATHS, &goPath{pth: abs})
		}
	}

	for _, pth = range strings.Split(os.Getenv("GOPATH"), string(os.PathListSeparator)) {
		if pth != "" {
			if _, ok = paths[pth]; !ok {
				GOPATHS = append(GOPATHS, &goPath{pth: pth})
				paths[pth] = nil
			}
		}
	}

	pth = build.Default.GOPATH
	if _, ok = paths[pth]; !ok {
		GOPATHS = append(GOPATHS, &goPath{pth: pth})
		paths[pth] = nil
	}

	for _, pth := range GOPATHS {
		pth.hasSrcDir = IsExistingDir(filepath.Join(pth.pth, "src"))
	}
}

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

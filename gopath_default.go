// +build !dev

package path_helpers

import (
	"os"
	"path/filepath"
	"strings"
)

var gopathroot string

func init() {
	gopathroot = os.Getenv("GITHUBCOM_MOISESPSENAGO_PATHHELPERS_GOPATH")
	if gopathroot == "" {
		gopathroot = filepath.Join("." + filepath.Base(os.Args[0]))
	}
	GOPATHS = append([]*goPath{{pth: gopathroot, hasSrcDir: true}}, GOPATHS...)
}

func cleanCalled(pth string, abs ...bool) string {
	if pth[0] == filepath.Separator {
		pth = pth[1:]
	}
	if len(abs) > 0 && abs[0] {
		if !strings.HasPrefix(pth, "src") {
			pth = "src" + string(filepath.Separator) + pth
		}
		pth = filepath.Join(gopathroot, pth)
	} else {
		pth = strings.TrimPrefix(pth, "src"+string(filepath.Separator))
	}

	return pth
}

func GetCalledFile(abs ...bool) string {
	return cleanCalled(GetCalledFileAbs(2), abs...)
}

func GetCalledDir(abs ...bool) string {
	return cleanCalled(filepath.Dir(GetCalledFileAbs(2)), abs...)
}

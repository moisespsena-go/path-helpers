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
		gopathroot = filepath.Join("."+filepath.Base(os.Args[0]), "fs")
	}
	GOPATHS = append(GOPATHS, &goPath{pth: gopathroot})
}

func cleanCalled(pth string, abs ...bool) string {
	pth = strings.TrimPrefix(pth, "src"+string(filepath.Separator))

	if len(abs) > 0 && abs[0] {
		pth = filepath.Join(gopathroot, pth)
	}
	return pth
}

func GetCalledFile(abs ...bool) string {
	return cleanCalled(GetCalledFileAbs(2), abs...)
}

func GetCalledDir(abs ...bool) string {
	return cleanCalled(filepath.Dir(GetCalledFileAbs(2)), abs...)
}

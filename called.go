package path_helpers

import (
	"errors"
	"runtime"
)

func GetCalledFileAbs(skip ...int) string {
	if len(skip) == 0 {
		skip = []int{1}
	}
	_, pth, _, ok := runtime.Caller(skip[0])
	if !ok {
		panic(errors.New("runtime.Caller: information unavailable."))
	}
	return pth
}

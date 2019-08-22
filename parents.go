package path_helpers

import "strings"

func Parents(sub, sep string, cb func(sub string) error) (err error) {
	if sub == sep {
		return
	}
	prefix := strings.HasPrefix(sub, sep)
	parts := strings.Split(strings.TrimPrefix(sub, sep), sep)
	l := len(parts)
	for i := l; i > 0; i-- {
		pth := strings.Join(parts[0:i], sep)
		if prefix {
			pth = sep + pth
		}
		if err = cb(pth); err != nil {
			return
		}
	}
	return
}

package urlpath

import (
	"path"
	"strings"
)

func ShiftPath(p string) (start, end string) {
	p = path.Clean("/" + p)
	//fmt.Println("After Clean: ", p)

	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	//fmt.Println("After Index: ", p[1:i])
	return p[1:i], p[i:]
}
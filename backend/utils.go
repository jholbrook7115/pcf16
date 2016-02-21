package backend

import "strings"

func clean(str string) string {
	str = strings.Join(strings.Split(str, "."), " ")
	str = strings.Join(strings.Split(str, ","), " ")
	return strings.ToLower(str)
}

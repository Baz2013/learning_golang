package lcof

import "strings"

func ReplaceSpace(s string) string {

	return strings.Replace(s, " ", "%20", -1)
}

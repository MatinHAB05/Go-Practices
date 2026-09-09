package filtercast

import "strings"

func ByName(Name string, prefix string) bool {
	return strings.HasPrefix(Name, prefix)
}

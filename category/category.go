package category

import (
	"strings"
)

func IsPopular(searchWord string) bool {
	if strings.Contains(searchWord, "_") {
		return true
	}
	return false
}

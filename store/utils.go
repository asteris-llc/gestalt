package store

import (
	"path"
	"strings"
)

func ensurePrefix(prefix, key string) string {
	if strings.HasPrefix(prefix, key) {
		return key
	}

	return path.Join(prefix, key)
}

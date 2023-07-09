package util

import "strings"

// StandardizeUri return standardized uri, remove trailing slash, replace ipfs link and so on
func StandardizeUri(s string) string {
	s = strings.TrimSpace(s)
	s = strings.Replace(s, "ipfs://", "https://consolelabs.mypinata.cloud/ipfs/", -1)
	return s
}

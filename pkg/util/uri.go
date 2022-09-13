package util

import "strings"

// StandardizeUri return standardized uri, remove trailing slash, replace ipfs link and so on
func StandardizeUri(s string) string {
	s = strings.TrimSpace(s)
	s = strings.Replace(s, "ipfs://", "https://consolelabs.mypinata.cloud/ipfs/", -1)
	return s
}

func UpvoteSourceNameAndUrl(source string) (name, url string) {
	switch source {
	case "topgg":
		return "Top.gg", "https://top.gg/bot/963123183131709480"
	case "discordbotlist":
		return "Discord Bot List", "https://discordbotlist.com/bots/mochi-bot"
	default:
		return source, ""
	}
}

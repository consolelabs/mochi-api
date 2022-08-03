package util

import "strings"

func GetTwitterRarityEmoji(rarity string) string {
	switch strings.ToLower(rarity) {
	case "common":
		return "⚪️"
	case "uncommon":
		return "🟢"
	case "rare":
		return "🔵"
	case "epic":
		return "🟣"
	case "legendary":
		return "🟠"
	case "mythic":
		return "🔴"
	default:
		return "⚪️"
	}
}

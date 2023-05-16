package util

import "strings"

func GetSuiAddressCollection(types string) string {
	typeArray := strings.Split(types, "::")
	if len(typeArray) > 0 {
		return typeArray[0]
	}
	return ""
}

func GetSymbolSuiCollection(slug string) string {
	slugArray := strings.Split(slug, "--")
	if len(slugArray) > 0 {
		temp := slugArray[0]
		if len(temp) > 10 {
			temp = temp[0:10]
		}
		symbol := strings.ReplaceAll(temp, "-", "")
		return symbol
	}
	temp := slug
	if len(temp) > 10 {
		temp = temp[0:10]
	}
	symbol := strings.ReplaceAll(temp, "-", "")
	return symbol
}

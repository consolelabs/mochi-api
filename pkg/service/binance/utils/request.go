package butils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

func signature(queryString, apiSecret string) string {
	h := hmac.New(sha256.New, []byte(apiSecret))
	h.Write([]byte(queryString))
	return hex.EncodeToString(h.Sum(nil))
}

func QueryString(q map[string]string, apiSecret string) string {
	var s []string
	for k, v := range q {
		s = append(s, fmt.Sprintf("%s=%s", k, v))
	}
	queryString := strings.Join(s, "&")
	signature := signature(queryString, apiSecret)
	return queryString + "&signature=" + signature
}

func FirstTimestamp() int64 {
	return 1498880255000
}

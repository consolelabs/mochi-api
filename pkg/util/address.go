package util

import (
	"encoding/hex"
	"strconv"
	"strings"

	"golang.org/x/crypto/sha3"
)

type ChainType string

const (
	EVM   ChainType = "evm"
	TERRA ChainType = "terra"
)

func GetChainTypeFromAddress(address string) ChainType {
	if strings.HasPrefix(address, string(TERRA)) {
		return TERRA
	}

	return EVM
}

func calculateKeccak256(addr []byte) []byte {
	hash := sha3.NewLegacyKeccak256()
	hash.Write(addr)
	return hash.Sum(nil)
}

func checksumByte(addr byte, hash byte) string {
	result := strconv.FormatUint(uint64(addr), 16)
	if hash >= 8 {
		return strings.ToUpper(result)
	} else {
		return result
	}
}

func ConvertToChecksumAddr(addrStr string) (string, error) {
	addrStr = addrStr[2:]
	addr, err := hex.DecodeString(addrStr)
	if err != nil {
		return "", err
	}
	hash := calculateKeccak256([]byte(strings.ToLower(addrStr)))

	result := "0x"

	for i, b := range addr {
		result += checksumByte(b>>4, hash[i]>>4)
		result += checksumByte(b&0xF, hash[i]&0xF)
	}

	return result, nil
}

func ShortenAddress(address string) string {
	if address == "" {
		return ""
	}

	return string(address[0:4]) + "..." + string(address[len(address)-4:])
}

package util

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func encodeSignMsg(data string) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256([]byte(msg))
}

func VerifySig(from, sigHex, msg string) error {
	fromAddr := common.HexToAddress(from)

	sig := hexutil.MustDecode(sigHex)

	/*
		https://github.com/ethereum/go-ethereum/blob/55599ee95d4151a2502465e0afc7c47bd1acba77/internal/ethapi/api.go#L442
		https://bitcoin.stackexchange.com/questions/38351/ecdsa-v-r-s-what-is-v
	*/
	switch sig[64] {
	case 0, 1:
	case 27, 28:
		sig[64] -= 27
	default:
		return fmt.Errorf("invalid ECDSA header V is %d not 0|1|27|28", sig[64])
	}

	pubKey, err := crypto.SigToPub(encodeSignMsg(msg), sig)
	if err != nil {
		return err
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey)

	if fromAddr != recoveredAddr {
		return fmt.Errorf("invalid signature")
	}

	return nil
}

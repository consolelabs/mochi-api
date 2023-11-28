package chain

import (
	"context"
	"math/big"

	"github.com/pkg/errors"
	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/liteclient"
	"github.com/xssnick/tonutils-go/ton"
	"github.com/xssnick/tonutils-go/ton/jetton"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/logger"
)

type TonClient struct {
	client *liteclient.ConnectionPool
	api    ton.APIClientWrapped
	logger logger.Logger
}

func NewTonClient(config *config.Config, l logger.Logger) (*TonClient, error) {
	client := liteclient.NewConnectionPool()

	// List ton lite servers
	configURL := "https://ton.org/global.config.json" // default by mainnet
	if url := config.TonConfigURL; url != "" {
		configURL = url
	}

	api := ton.NewAPIClient(client, ton.ProofCheckPolicyFast).WithRetry(2)

	if err := client.AddConnectionsFromConfigUrl(context.Background(), configURL); err != nil {
		return nil, errors.WithStack(err)
	}

	return &TonClient{
		client: client,
		api:    api,
		logger: l,
	}, nil
}

func (t *TonClient) GetJettonBalance(ownerAddr, jettonMasterAddr string) (*big.Int, error) {
	jettonMaster, err := address.ParseAddr(jettonMasterAddr)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	owner, err := address.ParseAddr(ownerAddr)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	jettonMasterContract := jetton.NewJettonMasterClient(t.api, jettonMaster)
	if jettonMasterContract == nil {
		return nil, errors.WithStack(errors.New("failed to create jetton master contract client"))
	}

	ownerJettonWallet, err := jettonMasterContract.GetJettonWallet(context.Background(), owner)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	jettonBalance, err := ownerJettonWallet.GetBalance(context.Background())
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return jettonBalance, nil
}

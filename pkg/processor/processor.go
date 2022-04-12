package processor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/defipod/mochi/pkg/config"
)

type processor struct {
	Config config.Config
}

func NewProcessor(cfg config.Config) Processor {
	return &processor{
		Config: cfg,
	}
}

type CreateUserTransaction struct {
	Dapp   string `json:"dapp"`
	Action string `json:"action"`
	Data   Data   `json:"data"`
}

type Data struct {
	TxHash             string  `json:"tx_hash,omitempty"`
	UserID             string  `json:"user_id,omitempty"`
	UserDiscordId      string  `json:"user_discord_id,omitempty"`
	UserWalletAddress  string  `json:"user_wallet_address,omitempty"`
	RecipientDiscordId string  `json:"recipient_discord_id,omitempty"`
	RecipientUserID    string  `json:"recipient_user_id,omitempty"`
	Guild              string  `json:"guild_id,omitempty"`
	Amount             float64 `json:"amount,omitempty"`
	Cryptocurrency     string  `json:"cryptocurrency,omitempty"`
	TokenPriceSymbol   string  `json:"token_price_symbol,omitempty"`
	PoolName           string  `json:"pool_name,omitempty"`
	PoolToken          string  `json:"pool_token,omitempty"`
	Link               string  `json:"link,omitempty"`
	TwitterAddress     string  `json:"twitter_address,omitempty"`
	NekoId             string  `json:"neko_id,omitempty"`
	NekoName           string  `json:"name,omitempty"`
	NekoNameBid        string  `json:"neko_name_bid,omitempty"`
	Time               string  `json:"time,omitempty"`
	StreakCount        int     `json:"streak_count,omitempty"`
	TotalCount         int     `json:"total_count,omitempty"`
	ContractAddress    string  `json:"contract_address,omitempty"`
}

func (p *processor) CreateUserTransaction(createUserTransactionRequest CreateUserTransaction) (int, error) {
	body, err := json.Marshal(createUserTransactionRequest)
	if err != nil {
		return 500, err
	}

	jsonBody := bytes.NewBuffer(body)

	var client = &http.Client{}
	request, err := http.NewRequest("POST", p.Config.ProcessorServerHost+"/user_transaction", jsonBody)
	if err != nil {
		return 500, err
	}

	request.Header.Add("Content-Type", "application/json")

	response, err := client.Do(request)

	if err != nil {
		return 500, err
	}

	if response.StatusCode != http.StatusOK {
		errBody := new(bytes.Buffer)
		_, err := errBody.ReadFrom(response.Body)
		if err != nil {
			return 500, fmt.Errorf("failed to read response body: %v", err)
		}

		return response.StatusCode, fmt.Errorf("failed to send user transaction api: %v", errBody.String())
	}

	defer response.Body.Close()

	return response.StatusCode, nil
}

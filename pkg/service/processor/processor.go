package processor

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/model"
)

type processor struct {
	config *config.Config
}

func NewProcessor(cfg *config.Config) Service {
	return &processor{
		config: cfg,
	}
}

func (p *processor) CreateUserTransaction(createUserTransactionRequest model.CreateUserTransaction) (*model.CreateUserTxResponse, error) {
	body, err := json.Marshal(createUserTransactionRequest)
	if err != nil {
		return nil, err
	}

	jsonBody := bytes.NewBuffer(body)

	var client = &http.Client{}
	request, err := http.NewRequest("POST", p.config.ProcessorServerHost+"/user_transaction", jsonBody)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", "application/json")

	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	resBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	res := &model.CreateUserTxResponse{}
	err = json.Unmarshal(resBody, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

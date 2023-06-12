package birdeye

import (
	"fmt"
	"net/http"

	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/util"
)

type birdeye struct {
	config *config.Config
	logger logger.Logger
}

func NewService(cfg *config.Config, l logger.Logger) Service {
	return &birdeye{
		config: cfg,
		logger: l,
	}
}

var (
	publicBirdeye = "https://public-api.birdeye.so"
)

func (b *birdeye) GetTokenPrice(address string) (*TokenPrice, error) {
	var res TokenPrice
	url := fmt.Sprintf("%s/public/price?address=%s", publicBirdeye, address)
	err := b.fetchBirdeyeData(url, &res)
	if err != nil {
		b.logger.Fields(logger.Fields{"url": url}).Error(err, "[birdeye.GetTokenPrice] b.fetchBirdeyeData() failed")
		return nil, err
	}
	return &res, nil
}

func (b *birdeye) fetchBirdeyeData(url string, v any) error {
	req := util.SendRequestQuery{
		URL:       url,
		ParseForm: &v,
		Headers:   map[string]string{"Content-Type": "application/json"},
	}

	statusCode, err := util.SendRequest(req)
	if err != nil || statusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch Birdeye data")
	}

	return nil
}

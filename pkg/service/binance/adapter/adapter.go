package bapdater

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/defipod/mochi/pkg/response"
	butils "github.com/defipod/mochi/pkg/service/binance/utils"
)

var (
	url = "https://api.binance.com"
)

func GetApiKeyPermission(apiKey, apiSecret string) (permission *response.BinanceApiKeyPermissionResponse, err error) {
	q := map[string]string{
		"timestamp": strconv.Itoa(int(time.Now().UnixMilli())),
	}
	queryString := butils.QueryString(q, apiSecret)

	// http request
	req, err := http.NewRequest("GET", url+"/sapi/v1/account/apiRestrictions?"+queryString, nil)
	if err != nil {
		return nil, err
	}

	resp, err := do(req, apiKey, 0)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// decode response json
	err = json.NewDecoder(resp.Body).Decode(&permission)
	if err != nil {
		return nil, err
	}

	return permission, nil
}

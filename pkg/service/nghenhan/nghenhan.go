package nghenhan

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/defipod/mochi/pkg/response"
	"github.com/defipod/mochi/pkg/util"
)

type Nghenhan struct {
	baseUrl string
}

func NewService() Service {
	return &Nghenhan{
		baseUrl: "https://cex.console.so/api/v1",
	}
}

func (n *Nghenhan) GetFiatHistoricalChart(base, target, interval string, limit int) (*response.NghenhanFiatHistoricalChartResponse, error) {
	tmpBase := base
	tmpTarg := target
	// current api only support USD base -> TODO(catngh): support other bases when nghenhan side done
	if base != "usd" {
		tmpTarg = tmpBase
		tmpBase = "usd"
	}
	url := n.baseUrl + fmt.Sprintf("/rate?base=%s&target=%s&interval=%s&limit=%v", tmpBase, tmpTarg, interval, limit)
	data := response.NghenhanFiatHistoricalChartResponse{}
	req := util.SendRequestQuery{
		URL:       url,
		ParseForm: &data,
		Headers:   map[string]string{"Content-Type": "application/json"},
	}
	statusCode, err := util.SendRequest(req)
	if err != nil || statusCode != http.StatusOK {
		return nil, fmt.Errorf("[nghenhan.GetFiatHistoricalChart] util.SendRequest() failed: %v", err)
	}

	// initial base is not USD, so we inverse all rate
	if base != "usd" {
		for i, _ := range data.Data {
			cPrice, _ := strconv.ParseFloat(data.Data[i].ClosePrice, 64)
			oPrice, _ := strconv.ParseFloat(data.Data[i].OpenPrice, 64)
			hPrice, _ := strconv.ParseFloat(data.Data[i].HighPrice, 64)
			lPrice, _ := strconv.ParseFloat(data.Data[i].LowPrice, 64)
			data.Data[i].ClosePrice = fmt.Sprintf("%f", 1/cPrice)
			data.Data[i].OpenPrice = fmt.Sprintf("%f", 1/oPrice)
			data.Data[i].HighPrice = fmt.Sprintf("%f", 1/hPrice)
			data.Data[i].LowPrice = fmt.Sprintf("%f", 1/lPrice)
		}
	}
	return &data, nil
}

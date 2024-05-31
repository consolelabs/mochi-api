package adapter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/defipod/mochi/pkg/util"
)

var (
	url = "https://api.dune.com"
)

func ExecuteQuery(apiKey string, queryId int64, param map[string]any) (res *ExecuteQueryResponse, err error) {
	endpoint := fmt.Sprintf("%s/api/v1/query/%v/execute", url, queryId)

	payload, err := json.Marshal(struct {
		QueryParameters map[string]any `json:"query_parameters"`
	}{QueryParameters: param})
	if err != nil {
		return nil, err
	}

	// http request
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	resp, err := do3(req, apiKey)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	return res, nil
}

func GetExecutionResult(apiKey, executionId string, limit, offset int64) (res *GetExecutionResultResponse, err error) {
	endpoint := fmt.Sprintf("%s/api/v1/execution/%s/result", url, executionId)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	if limit != 0 {
		q.Add("limit", fmt.Sprintf("%v", limit))
	}
	if offset != 0 {
		q.Add("offset", fmt.Sprintf("%v", offset))
	}
	req.URL.RawQuery = q.Encode()
	reqQuery := util.SendRequestQuery{
		URL:      req.URL.String(),
		Response: &res,
		Headers: map[string]string{
			"Content-Type":   "application/json",
			"X-DUNE-API-KEY": apiKey,
		},
	}
	_, err = util.SendRequest(reqQuery)
	if err != nil {
		return nil, err
	}
	fmt.Printf("res: %v\n", res)

	// resp, err := do3(req, apiKey)
	// if err != nil {
	// 	return nil, err
	// }
	// defer resp.Body.Close()
	// // defer res.Body.Close()
	// body, _ := ioutil.ReadAll(resp.Body)

	// fmt.Println(res)
	// fmt.Println(string(body))

	// // bodyBytes, err := io.ReadAll(resp.Body)
	// // if err != nil {
	// // 	return nil, err
	// // }

	// if err := json.Unmarshal(body, &res); err != nil {
	// 	return nil, err
	// }

	return res, nil
}

func GetExecutionStatus(apiKey, executionId string) (res *GetExecutionResultResponse, err error) {
	endpoint := fmt.Sprintf("%s/api/v1/execution/%s/status", url, executionId)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := do3(req, apiKey)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// decode response json
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "https://api.geckoterminal.com/api/v2/networks/eth/pools/0xf9d14ac21fa7c235fb6353d1bb23d6f213f79b8b"

	cl := http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("content-type", "application/json")
	req.Header.Add("user-agent", "PostmanRuntime/7.32.3")

	resp, err := cl.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	bodyStr := string(body)

	fmt.Println(bodyStr)
}

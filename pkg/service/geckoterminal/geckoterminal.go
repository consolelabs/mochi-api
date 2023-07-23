package geckoterminal

import (
	"encoding/json"
	"fmt"

	"github.com/defipod/mochi/pkg/config"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/stealth"
)

type GeckoTerminal struct {
	chromeHost  string
	searchApi   string
	getPoolApi  string
	getPoolPage string
}

func NewService(cfg *config.Config) Service {
	return &GeckoTerminal{
		chromeHost:  cfg.ChromeHost,
		searchApi:   "https://app.geckoterminal.com/api/p1/search?query=%s",
		getPoolApi:  "https://api.geckoterminal.com/api/v2/networks/%s/pools/%s",
		getPoolPage: "https://www.geckoterminal.com/%s/pools/%s",
	}
}

func (g *GeckoTerminal) Search(query string) (*Search, error) {
	browser := rod.New().ControlURL(launcher.MustResolveURL(g.chromeHost)).MustConnect()
	defer browser.MustClose()
	page := stealth.MustPage(browser).MustNavigate(fmt.Sprintf(g.searchApi, query))

	data := page.MustElement("body").MustText()

	var search *Search
	if err := json.Unmarshal([]byte(data), &search); err != nil {
		return nil, err
	}

	return search, nil
}

func (g *GeckoTerminal) GetPool(network, pool string) (*Pool, error) {
	var poolResp *Pool
	browser := rod.New().ControlURL(launcher.MustResolveURL(g.chromeHost)).MustConnect()
	defer browser.MustClose()

	page := stealth.MustPage(browser).MustNavigate(fmt.Sprintf(g.getPoolApi, network, pool))
	data := page.MustElement("body").MustText()

	if err := json.Unmarshal([]byte(data), &poolResp); err != nil {
		return nil, err
	}

	return poolResp, nil
}

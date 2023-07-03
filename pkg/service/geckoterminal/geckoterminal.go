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
		getPoolApi:  "https://app.geckoterminal.com/api/p1/%s/pools/%s",
		getPoolPage: "https://www.geckoterminal.com/%s/pools/%s",
	}
}

func (g *GeckoTerminal) Search(query string) (*GeckoTerminalSearch, error) {
	browser := rod.New().ControlURL(launcher.MustResolveURL(g.chromeHost)).MustConnect()
	defer browser.MustClose()
	page := stealth.MustPage(browser).MustNavigate(fmt.Sprintf(g.searchApi, query))

	data := page.MustElement("body").MustText()

	var geckoTerminalResponse *GeckoTerminalSearch
	if err := json.Unmarshal([]byte(data), &geckoTerminalResponse); err != nil {
		return nil, err
	}

	return geckoTerminalResponse, nil
}

func (g *GeckoTerminal) GetPool(network, pool string) (*GeckoTerminalGetPool, error) {
	browser := rod.New().ControlURL(launcher.MustResolveURL(g.chromeHost)).MustConnect()
	defer browser.MustClose()

	page := stealth.MustPage(browser).MustNavigate(fmt.Sprintf(g.getPoolApi, network, pool))
	data := page.MustElement("body").MustText()

	var geckoTerminalResponse *GeckoTerminalGetPool
	if err := json.Unmarshal([]byte(data), &geckoTerminalResponse); err != nil {
		return nil, err
	}

	return geckoTerminalResponse, nil
}

func (g *GeckoTerminal) ScrapePool(network, pool string) (*ScrapePool, error) {
	data := &ScrapePool{}
	browser := rod.New().ControlURL(launcher.MustResolveURL(g.chromeHost)).MustConnect()
	defer browser.MustClose()

	page := stealth.MustPage(browser).MustNavigate(fmt.Sprintf(g.getPoolPage, network, pool))

	els1 := page.MustElements("[class='flex flex-shrink flex-grow flex-col md:col-span-2']")

	for i, el := range els1 {
		if el.MustElement("td") == nil {
			continue
		}

		switch i {
		case 0:
			data.Volume24h = el.MustElement("td").MustText()
		case 1:
			data.Liquidity = el.MustElement("td").MustText()
		case 2:
			data.FullyDilutedValuation = el.MustElement("td").MustText()
		case 3:
			data.MarketCap = el.MustElement("td").MustText()
		}
	}

	els2 := page.MustElement("[id='pool-price-display']")
	if els2 != nil {
		data.PriceInUSD = els2.MustElement("span").MustText()
	}

	return data, nil
}

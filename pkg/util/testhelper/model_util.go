package testhelper

import "github.com/defipod/mochi/pkg/response"

func WatchlistCoingeckoDefaultTickers() []response.CoinMarketItemData {
	return []response.CoinMarketItemData{
		{
			ID:           "bitcoin",
			Name:         "Bitcoin",
			Symbol:       "btc",
			CurrentPrice: 19238.62,
			Image:        "https://assets.coingecko.com/coins/images/1/large/bitcoin.png?1547033579",
			SparkLineIn7d: struct {
				Price []float64 "json:\"price\""
			}{
				[]float64{
					18861.78098812819, 19087.709424228065, 19165.140362664395,
				},
			},
			PriceChangePercentage24h:          0.33071,
			PriceChangePercentage7dInCurrency: 2.281861441149524,
			IsPair:                            false,
		},
		{
			ID:           "ethereum",
			Name:         "Ethereum",
			Symbol:       "eth",
			CurrentPrice: 1302.07,
			Image:        "https://assets.coingecko.com/coins/images/279/large/ethereum.png?1595348880",
			SparkLineIn7d: struct {
				Price []float64 "json:\"price\""
			}{
				[]float64{
					1296.873383685728, 1317.4021123270202, 1321.6466884030708,
				},
			},
			PriceChangePercentage24h:          0.30844,
			PriceChangePercentage7dInCurrency: 0.6174474468422891,
			IsPair:                            false,
		},
		{
			ID:           "binancecoin",
			Name:         "BNB",
			Symbol:       "bnb",
			CurrentPrice: 285.96,
			Image:        "https://assets.coingecko.com/coins/images/825/large/bnb-icon2_2x.png?1644979850",
			SparkLineIn7d: struct {
				Price []float64 "json:\"price\""
			}{
				[]float64{
					271.8725401305167, 274.3169987349479, 274.52923955490246,
				},
			},
			PriceChangePercentage24h:          1.16908,
			PriceChangePercentage7dInCurrency: 4.332244779295762,
			IsPair:                            false,
		},
		{
			ID:           "solana",
			Name:         "Solana",
			Symbol:       "sol",
			CurrentPrice: 32.68,
			Image:        "https://assets.coingecko.com/coins/images/4128/large/solana.png?1640133422",
			SparkLineIn7d: struct {
				Price []float64 "json:\"price\""
			}{
				[]float64{
					32.55379876249857, 32.853415888089245, 33.105954022206525,
				},
			},
			PriceChangePercentage24h:          0.41472,
			PriceChangePercentage7dInCurrency: 1.103904029633381,
			IsPair:                            false,
		},
		{
			ID:           "matic-network",
			Name:         "Polygon",
			Symbol:       "matic",
			CurrentPrice: 0.775881,
			Image:        "https://assets.coingecko.com/coins/images/4713/large/matic-token-icon.png?1624446912",
			SparkLineIn7d: struct {
				Price []float64 "json:\"price\""
			}{
				[]float64{
					0.7407198851382445, 0.7502069119119806, 0.7559222981580057,
				},
			},
			PriceChangePercentage24h:          0.81771,
			PriceChangePercentage7dInCurrency: 4.725445455130517,
			IsPair:                            false,
		},
		{
			ID:           "avalanche-2",
			Name:         "Avalanche",
			Symbol:       "avax",
			CurrentPrice: 16.86,
			Image:        "https://assets.coingecko.com/coins/images/12559/large/coin-round-red.png?1604021818",
			SparkLineIn7d: struct {
				Price []float64 "json:\"price\""
			}{
				[]float64{
					17.205317839234308, 17.36749528376322, 17.43106660592316,
				},
			},
			PriceChangePercentage24h:          0.35013,
			PriceChangePercentage7dInCurrency: -2.6628978539310992,
			IsPair:                            false,
		},
		{
			ID:           "internet-computer",
			Name:         "Internet Computer",
			Symbol:       "icp",
			CurrentPrice: 6,
			Image:        "https://assets.coingecko.com/coins/images/14495/large/Internet_Computer_logo.png?1620703073",
			SparkLineIn7d: struct {
				Price []float64 "json:\"price\""
			}{
				[]float64{
					6.168892847271933, 6.155438275738423, 6.131422137275525,
				},
			},
			PriceChangePercentage24h:          -0.39209,
			PriceChangePercentage7dInCurrency: 1.6237690910330411,
			IsPair:                            false,
		},
		{
			ID:           "fantom",
			Name:         "Fantom",
			Symbol:       "ftm",
			CurrentPrice: 0.221302,
			Image:        "https://assets.coingecko.com/coins/images/4001/large/Fantom.png?1558015016",
			SparkLineIn7d: struct {
				Price []float64 "json:\"price\""
			}{
				[]float64{
					0.22630507054947152, 0.22775235232649785, 0.22848573456673352,
				},
			},
			PriceChangePercentage24h:          -0.01195,
			PriceChangePercentage7dInCurrency: -1.428104703485673,
			IsPair:                            false,
		},
	}
}

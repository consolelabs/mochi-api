
-- +migrate Up
INSERT INTO tokens (id,address,symbol,chain_id,decimals,discord_bot_supported,coin_gecko_id,"name",guild_default,is_native) VALUES
    (1,'0x321162Cd933E2Be498Cd2267a90534A804051b11','BTC',250,8,true,'bitcoin','Bitcoin',false,false),
    (2,'0x74b23882a30290451A17c44f4F05243b6b58C76d','ETH',250,18,true,'ethereum','Ethereum',false,false),
    (3,'0xdAC17F958D2ee523a2206206994597C13D831ec7','USDT',1,6,true,'tether','Tether',false,false),
    (4,'0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48','USDC',1,6,true,'usd-coin','USD Coin',false,false),
    (8,'0x3EE2200Efb3400fAbB9AacF31297cBdD1d435D47','ADA',56,18,true,'binance-peg-cardano','Binance-Peg Cardano',false,false),
    (5,'0xB8c77482e45F1F44dE1745F52C74426C631bDD52','BNB',1,18,true,'binancecoin','BNB',false,false),
    (6,'0x1D2F0da169ceB9fC7B3144628dB156f3F6c60dBE','XRP',56,18,true,'ripple','Binance-Peg XRP',false,false),
    (7,'0x4Fabb145d64652a948d72533023f6E7A623C7C53','BUSD',1,18,true,'binance-usd','Binance USD',false,false),
    (9,'0x7083609fCE4d1d8Dc0C979AAb8c869Ea2C873402','DOT',56,18,true,'binance-peg-polkadot','Binance-Peg Polkadot',false,false),
    (10,'0xbA2aE424d960c26247Dd6c32edC70B295c744C43','DOGE',56,8,true,'binance-peg-dogecoin','Binance-Peg Dogecoin',false,false),
    (11,'0x6B175474E89094C44Da98b954EedeAC495271d0F','DAI',1,18,true,'dai','Dai',false,false),
    (12,'0x7D1AfA7B718fb893dB30A3aBc0Cfc608AaCfeBB0','MATIC',1,18,true,'matic-network','Matic',false,false),
    (13,'0x1CE0c2827e2eF14D5C4f29a091d735A204794041','AVAX',56,18,true,'binance-peg-avalanche','Binance-Peg Avalanche',false,false),
    (14,'0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984','UNI',1,18,true,'uniswap','Uniswap',false,false),
    (15,'0x95aD61b0a150d79219dCF64E1E6Cc01f0B64C4cE','SHIB',1,18,true,'shiba-inu','SHIBA INU',false,false),
    (16,'0xE1Be5D3f34e89dE342Ee97E6e90D405884dA6c67','TRX',1,6,true,'tron','TRON',false,false),
    (17,'0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599','WBTC',1,8,true,'wrapped-bitcoin','Wrapped BTC',false,false),
    (18,'0x3d6545b08693daE087E957cb1180ee38B9e3c25E','ETC',56,18,true,'ethereum-classic','Binance-Peg Ethereum Classic',false,false),
    (19,'0x2AF5D2aD76741191D15Dfe7bF6aC92d4Bd912Ca3','LEO',1,18,true,'leo-token','Bitfinex LEO',false,false),
    (20,'0x4338665CBB7B2485A8855A139b75D5e34AB0DB94','LTC',56,18,true,'binance-peg-litecoin','Binance-Peg Litecoin',false,false),
    (21,'0x50D1c9771902476076eCFc8B2A83Ad6b9355a4c9','FTT',1,18,true,'ftx-token','FTX',false,false),
    (22,'0xA0b73E1Ff0B80914AB6fe0444E65848C4C34450b','CRO',1,8,true,'crypto-com-chain','Cronos Coin',false,false),
    (23,'0x514910771AF9Ca656af840dff83E8264EcF986CA','LINK',1,18,true,'chainlink','ChainLink',false,false),
    (24,'0x1Fa4a73a3F0133f0025378af00236f3aBDEE5D63','NEAR',56,18,true,'near','Binance-Peg NEAR Protocol',false,false),
    (25,'0x0Eb3a705fc54725037CC9e008bDede697f62F335','ATOM',56,18,true,'cosmos','Binance-Peg Cosmos',false,false),
    (26,'0xD893925F2035C40663F13CCDcB42Dd0a1C72a944','XLM',1,18,true,'stellar','Stellar',false,false),
    (27,'0x991170CDe1B4E90907A7C0515123C1B18D635107','XMR',56,18,true,'monero','Monero',false,false),
    (28,'0x8fF795a6F4D97E7887C79beA79aba5cc76444aDf','BCH',56,18,true,'binance-peg-bitcoin-cash','Binance-Peg Bitcoin Cash',false,false),
    (29,'0x4d224452801ACEd8B2F0aebE155379bb5D594381','APE',1,18,true,'apecoin','ApeCoin',false,false),
    (30,'0x4E15361FD6b4BB609Fa63C81A2be19d873717870','FTM',1,18,true,'fantom','Fantom',false,false);

-- +migrate Down
DELETE FROM tokens WHERE id <= 30;

package config

import (
	"strings"
	"time"

	"github.com/spf13/viper"
)

const (
	publicCertsTTL = 24
)

// Config contain configuration of db for migrator
// config var < env < command flag
type Config struct {
	ServiceName    string
	BaseURL        string
	Port           string
	Env            string
	AllowedOrigins string
	DBHost         string
	DBPort         string
	DBUser         string
	DBName         string
	DBPass         string
	DBSSLMode      string

	DiscordLogWebhook       string
	InDiscordWalletMnemonic string

	MochiBotSecret string

	JWTSecret              []byte
	JWTAccessTokenLifeSpan time.Duration

	FantomRPC        string
	FantomScan       string
	FantomScanAPIKey string

	EthereumRPC        string
	EthereumScan       string
	EthereumScanAPIKey string

	BscRPC        string
	BscScan       string
	BscScanAPIKey string

	CovalentAPIKey string

	DiscordToken string

	RedisURL string

	MochiLogChannelID  string
	MochiSaleChannelID string

	MoralisXApiKey string

	IndexerServerHost string

	RpcUrl RpcUrl

	MarketplaceBaseUrl MarketplaceBaseUrl

	MarketplaceApiKey        MarketplaceApiKey
	TwitterAccessToken       string
	TwitterAccessTokenSecret string
	TwitterConsumerKey       string
	TwitterConsumerSecret    string
}

type MarketplaceBaseUrl struct {
	Opensea  string
	Quixotic string
	Painswap string
}

type MarketplaceApiKey struct {
	Opensea  string
	Quixotic string
}
type RpcUrl struct {
	Eth string
	Ftm string
	Opt string
}

// GetCORS in config
func (c *Config) GetCORS() []string {
	cors := strings.Split(c.AllowedOrigins, ";")
	rs := []string{}
	for idx := range cors {
		itm := cors[idx]
		if strings.TrimSpace(itm) != "" {
			rs = append(rs, itm)
		}
	}
	return rs
}

// Loader load config from reader into Viper
type Loader interface {
	Load(viper.Viper) (*viper.Viper, error)
}

// generateConfigFromViper generate config from viper data
func generateConfigFromViper(v *viper.Viper) Config {
	tokenTTLInDay := v.GetInt("ACCESS_TOKEN_TTL")
	if tokenTTLInDay == 0 {
		tokenTTLInDay = 7
	}

	return Config{
		Port:        v.GetString("PORT"),
		BaseURL:     v.GetString("BASE_URL"),
		ServiceName: v.GetString("SERVICE_NAME"),
		Env:         v.GetString("ENV"),

		AllowedOrigins: v.GetString("ALLOWED_ORIGINS"),

		DBHost:    v.GetString("DB_HOST"),
		DBPort:    v.GetString("DB_PORT"),
		DBUser:    v.GetString("DB_USER"),
		DBName:    v.GetString("DB_NAME"),
		DBPass:    v.GetString("DB_PASS"),
		DBSSLMode: v.GetString("DB_SSL_MODE"),

		MochiBotSecret: v.GetString("MOCHI_BOT_SECRET"),

		JWTSecret:              []byte(v.GetString("JWT_SECRET")),
		JWTAccessTokenLifeSpan: time.Hour * 24 * time.Duration(tokenTTLInDay), // 7 days

		FantomRPC:        v.GetString("FANTOM_RPC"),
		FantomScan:       v.GetString("FANTOM_SCAN"),
		FantomScanAPIKey: v.GetString("FANTOM_SCAN_API_KEY"),

		EthereumRPC:        v.GetString("ETHEREUM_RPC"),
		EthereumScan:       v.GetString("ETHEREUM_SCAN"),
		EthereumScanAPIKey: v.GetString("ETHEREUM_SCAN_API_KEY"),

		BscRPC:        v.GetString("BSC_RPC"),
		BscScan:       v.GetString("BSC_SCAN"),
		BscScanAPIKey: v.GetString("BSC_SCAN_API_KEY"),

		CovalentAPIKey: v.GetString("COVALENT_API_KEY"),

		DiscordToken: v.GetString("DISCORD_TOKEN"),

		InDiscordWalletMnemonic: v.GetString("IN_DISCORD_WALLET_MNEMONIC"),
		RedisURL:                v.GetString("REDIS_URL"),

		MochiLogChannelID:  v.GetString("MOCHI_LOG_CHANNEL_ID"),
		MochiSaleChannelID: v.GetString("MOCHI_SALE_CHANNEL_ID"),

		MoralisXApiKey: v.GetString("MORALIS_X_API_KEY"),

		IndexerServerHost: v.GetString("INDEXER_SERVER_HOST"),

		RpcUrl: RpcUrl{
			Eth: v.GetString("ETH_RPC"),
			Ftm: v.GetString("FTM_RPC"),
			Opt: v.GetString("OPTIMISM_RPC"),
		},

		MarketplaceBaseUrl: MarketplaceBaseUrl{
			Opensea:  v.GetString("OPENSEA_BASE_URL"),
			Quixotic: v.GetString("QUIXOTIC_BASE_URL"),
			Painswap: v.GetString("PAINTSWAP_BASE_URL"),
		},
		MarketplaceApiKey: MarketplaceApiKey{
			Opensea:  v.GetString("OPENSEA_API_KEY"),
			Quixotic: v.GetString("QUIXOTIC_API_KEY"),
		},
		TwitterAccessToken:       v.GetString("TWITTER_ACCESS_TOKEN"),
		TwitterAccessTokenSecret: v.GetString("TWITTER_ACCESS_TOKEN_SECRET"),
		TwitterConsumerKey:       v.GetString("TWITTER_CONSUMER_KEY"),
		TwitterConsumerSecret:    v.GetString("TWITTER_CONSUMER_SECRET"),
	}
}

// DefaultConfigLoaders is default loader list
func DefaultConfigLoaders() []Loader {
	loaders := []Loader{}
	fileLoader := NewFileLoader(".env", ".")
	loaders = append(loaders, fileLoader)
	loaders = append(loaders, NewENVLoader())

	return loaders
}

// LoadConfig load config from loader list
func LoadConfig(loaders []Loader) Config {
	v := viper.New()
	v.SetDefault("PORT", "8080")
	v.SetDefault("ENV", "local")
	v.SetDefault("FTM_RPC", "https://rpc.ankr.com/fantom")
	v.SetDefault("ETH_RPC", "https://rpc.ankr.com/eth")
	v.SetDefault("OPTIMISM_RPC", "https://rpc.ankr.com/optimism")
	v.SetDefault("OPENSEA_BASE_URL", "https://api.opensea.io")
	v.SetDefault("PAINTSWAP_BASE_URL", "https://api.paintswap.finance")
	v.SetDefault("QUIXOTIC_BASE_URL", "https://api.quixotic.io")

	for idx := range loaders {
		newV, err := loaders[idx].Load(*v)

		if err == nil {
			v = newV
		}
	}
	return generateConfigFromViper(v)
}

// GetShutdownTimeout get shutdown time out
func (c *Config) GetShutdownTimeout() time.Duration {
	return 10 * time.Second
}

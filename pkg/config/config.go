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
	ServiceName string
	BaseURL     string
	Port        string
	Env         string
	Debug       bool

	AllowedOrigins string
	DBHost         string
	DBPort         string
	DBUser         string
	DBName         string
	DBPass         string
	DBSSLMode      string
	DBReadHosts    []string

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

	CovalentBaseUrl string
	CovalentAPIKeys []string

	DiscordToken string

	RedisURL string

	MochiGuildID           string
	MochiLogChannelID      string
	MochiSaleChannelID     string
	MochiActivityChannelID string
	MochiFeedbackChannelID string

	MoralisXApiKey string

	IndexerServerHost string

	PodtownServerHost string

	RpcUrl RpcUrl

	MarketplaceBaseUrl MarketplaceBaseUrl

	MarketplaceApiKey        MarketplaceApiKey
	TwitterAccessToken       string
	TwitterAccessTokenSecret string
	TwitterConsumerKey       string
	TwitterConsumerSecret    string

	GoogleCloudBucketName     string
	GoogleCloudProjectID      string
	GoogleCloudServiceAccount string

	AppleKeyID   string
	AppleTeamID  string
	AppleAuthKey string

	ProcessorServerHost string

	BlockChainAPIKeyID     string
	BlockChainAPISecretKey string

	CoinGeckoAPIKey string

	CentralizedWalletPrivateKey string
	CentralizedWalletAddress    string

	SolanaCentralizedWalletPrivateKey string
	SolanaPKSecretKey                 string

	APILayerAPIKey string

	Kafka   Kafka
	Solscan Solscan
}

type MarketplaceBaseUrl struct {
	Opensea       string
	Quixotic      string
	Painswap      string
	BluemoveAptos string
	BluemoveSui   string
}

type Solscan struct {
	Token string
}

type MarketplaceApiKey struct {
	Opensea  string
	Quixotic string
}
type RpcUrl struct {
	Eth      string
	Ftm      string
	Opt      string
	Bsc      string
	Polygon  string
	Arbitrum string
	Okc      string
	Onus     string
}

type Kafka struct {
	Brokers string
	Topic   string
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
		Debug:       v.GetBool("DEBUG") || false,

		AllowedOrigins: v.GetString("ALLOWED_ORIGINS"),

		DBHost:      v.GetString("DB_HOST"),
		DBPort:      v.GetString("DB_PORT"),
		DBUser:      v.GetString("DB_USER"),
		DBName:      v.GetString("DB_NAME"),
		DBPass:      v.GetString("DB_PASS"),
		DBSSLMode:   v.GetString("DB_SSL_MODE"),
		DBReadHosts: strings.Split(v.GetString("DB_READ_HOSTS"), ","),

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

		CovalentBaseUrl: v.GetString("COVALENT_BASE_URL"),
		CovalentAPIKeys: strings.Split(v.GetString("COVALENT_API_KEYS"), ","),

		DiscordToken: v.GetString("DISCORD_TOKEN"),

		InDiscordWalletMnemonic: v.GetString("IN_DISCORD_WALLET_MNEMONIC"),
		RedisURL:                v.GetString("REDIS_URL"),

		MochiGuildID:           v.GetString("MOCHI_GUILD_ID"),
		MochiLogChannelID:      v.GetString("MOCHI_LOG_CHANNEL_ID"),
		MochiSaleChannelID:     v.GetString("MOCHI_SALE_CHANNEL_ID"),
		MochiActivityChannelID: v.GetString("MOCHI_ACTIVITY_CHANNEL_ID"),
		MochiFeedbackChannelID: v.GetString("MOCHI_FEEDBACK_CHANNEL_ID"),

		MoralisXApiKey: v.GetString("MORALIS_X_API_KEY"),

		IndexerServerHost: v.GetString("INDEXER_SERVER_HOST"),

		PodtownServerHost: v.GetString("PODTOWN_SERVER_HOST"),

		RpcUrl: RpcUrl{
			Eth:      v.GetString("ETH_RPC"),
			Ftm:      v.GetString("FTM_RPC"),
			Opt:      v.GetString("OPTIMISM_RPC"),
			Bsc:      v.GetString("BSC_RPC"),
			Arbitrum: v.GetString("ARBITRUM_RPC"),
			Polygon:  v.GetString("POLYGON_RPC"),
			Okc:      v.GetString("OKC_RPC"),
			Onus:     v.GetString("ONUS_RPC"),
		},

		MarketplaceBaseUrl: MarketplaceBaseUrl{
			Opensea:       v.GetString("OPENSEA_BASE_URL"),
			Quixotic:      v.GetString("QUIXOTIC_BASE_URL"),
			Painswap:      v.GetString("PAINTSWAP_BASE_URL"),
			BluemoveAptos: v.GetString("BLUEMOVE_APTOS_BASE_URL"),
			BluemoveSui:   v.GetString("BLUEMOVE_SUI_BASE_URL"),
		},
		MarketplaceApiKey: MarketplaceApiKey{
			Opensea:  v.GetString("OPENSEA_API_KEY"),
			Quixotic: v.GetString("QUIXOTIC_API_KEY"),
		},
		TwitterAccessToken:       v.GetString("TWITTER_ACCESS_TOKEN"),
		TwitterAccessTokenSecret: v.GetString("TWITTER_ACCESS_TOKEN_SECRET"),
		TwitterConsumerKey:       v.GetString("TWITTER_CONSUMER_KEY"),
		TwitterConsumerSecret:    v.GetString("TWITTER_CONSUMER_SECRET"),

		GoogleCloudBucketName:     v.GetString("GOOGLE_CLOUD_BUCKET_NAME"),
		GoogleCloudProjectID:      v.GetString("GOOGLE_CLOUD_PROJECT_ID"),
		GoogleCloudServiceAccount: v.GetString("GCP_SERVICE_ACCOUNT"),

		AppleKeyID:   v.GetString("APPLE_KEY_ID"),
		AppleTeamID:  v.GetString("APPLE_TEAM_ID"),
		AppleAuthKey: v.GetString("APPLE_AUTH_KEY"),

		ProcessorServerHost:    v.GetString("PROCESSOR_SERVER_HOST"),
		BlockChainAPIKeyID:     v.GetString("BLOCKCHAIN_API_KEY_ID"),
		BlockChainAPISecretKey: v.GetString("BLOCKCHAIN_API_SECRET_KEY"),

		CoinGeckoAPIKey: v.GetString("COINGECKO_API_KEY"),

		CentralizedWalletPrivateKey: v.GetString("CENTRALIZED_WALLET_PRIVATE_KEY"),
		CentralizedWalletAddress:    v.GetString("CENTRALIZED_WALLET_ADDRESS"),

		SolanaCentralizedWalletPrivateKey: v.GetString("SOLANA_CENTRALIZED_WALLET_PK"),
		SolanaPKSecretKey:                 v.GetString("SOLANA_PK_SECRET_KEY"),

		APILayerAPIKey: v.GetString("API_LAYER_API_KEY"),
		Kafka: Kafka{
			Brokers: v.GetString("KAFKA_BROKERS"),
			Topic:   v.GetString("KAFKA_TOPIC"),
		},
		Solscan: Solscan{
			Token: v.GetString("SOLSCAN_TOKEN"),
		},
	}
}

// Config for testing
func LoadTestConfig() Config {
	cfg := Config{
		DBUser: "postgres",
		DBPass: "postgres",
		DBHost: "localhost",
		DBPort: "25433",
		DBName: "mochi_local_test",

		InDiscordWalletMnemonic: "holiday frequent toy bachelor auto use style result recycle crumble glue blouse",
		FantomRPC:               "sample",
		FantomScan:              "sample",
		FantomScanAPIKey:        "sample",

		EthereumRPC:        "sample",
		EthereumScan:       "sample",
		EthereumScanAPIKey: "sample",

		BscRPC:        "sample",
		BscScan:       "sample",
		BscScanAPIKey: "sample",

		DiscordToken: "sample",

		RedisURL: "redis://localhost:6379/0",
	}
	return cfg
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
	v.SetDefault("BSC_RPC", "https://rpc.ankr.com/bsc")
	v.SetDefault("ARBITRUM_RPC", "https://rpc.ankr.com/arbitrum")
	v.SetDefault("POLYGON_RPC", "https://rpc.ankr.com/polygon")
	v.SetDefault("OKC_RPC", "https://exchainrpc.okex.org")
	v.SetDefault("ONUS_RPC", "https://rpc.onuschain.io")
	v.SetDefault("OPENSEA_BASE_URL", "https://api.opensea.io")
	v.SetDefault("PAINTSWAP_BASE_URL", "https://api.paintswap.finance")
	v.SetDefault("QUIXOTIC_BASE_URL", "https://api.quixotic.io")
	v.SetDefault("BLUEMOVE_APTOS_BASE_URL", "https://3rd.console.so/bluemove/api")
	v.SetDefault("BLUEMOVE_SUI_BASE_URL", "https://3rd.console.so/bluemove/api")
	v.SetDefault("COVALENT_BASE_URL", "https://api.covalenthq.com/v1")
	v.SetDefault("CENTRALIZED_WALLET_ADDRESS", "0x4ec16127e879464bef6ab310084facec1e4fe465")
	v.SetDefault("SOLSCAN_TOKEN", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkQXQiOjE2NzU3NzcyODYxMjgsImVtYWlsIjoibmdvdHJvbmdraG9pMTEyQGdtYWlsLmNvbSIsImFjdGlvbiI6InRva2VuLWFwaSIsImlhdCI6MTY3NTc3NzI4Nn0.DCT8Fh8j9uWVpnQSMnq0uuzqeBngNLxc4r8a1Aa2C4Q")

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

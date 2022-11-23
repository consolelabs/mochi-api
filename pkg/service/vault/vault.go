package vault

import (
	"fmt"
	"strconv"
	"time"

	"github.com/defipod/mochi/pkg/config"
	vault "github.com/hashicorp/vault/api"
)

type Vault struct {
	data map[string]interface{}
}

func New(cfg *config.Config) (VaultService, error) {
	config := vault.DefaultConfig()
	config.Address = cfg.Vault.Address

	client, err := vault.NewClient(config)
	if err != nil {
		return nil, fmt.Errorf("unable to initialize Vault client: %v", err)
	}

	client.SetToken(cfg.Vault.Token)

	secret, err := client.Logical().Read(cfg.Vault.Path)
	if err != nil {
		return nil, fmt.Errorf("unable to read secret: %v", err)
	}
	if secret == nil {
		return nil, fmt.Errorf("unable to read secret in path")
	}

	data, ok := secret.Data["data"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("unable to read secret data")
	}

	return &Vault{
		data: data,
	}, nil
}

func (v *Vault) LoadConfig() *config.Config {
	tokenTTLInDay := v.GetInt("ACCESS_TOKEN_TTL")
	if tokenTTLInDay == 0 {
		tokenTTLInDay = 7
	}

	return &config.Config{
		Port:        v.GetString("PORT"),
		BaseURL:     v.GetString("BASE_URL"),
		ServiceName: v.GetString("SERVICE_NAME"),
		Env:         v.GetString("ENV"),
		Debug:       v.GetBool("DEBUG") || false,

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

		MochiGuildID:           v.GetString("MOCHI_GUILD_ID"),
		MochiLogChannelID:      v.GetString("MOCHI_LOG_CHANNEL_ID"),
		MochiSaleChannelID:     v.GetString("MOCHI_SALE_CHANNEL_ID"),
		MochiActivityChannelID: v.GetString("MOCHI_ACTIVITY_CHANNEL_ID"),
		MochiFeedbackChannelID: v.GetString("MOCHI_FEEDBACK_CHANNEL_ID"),

		MoralisXApiKey: v.GetString("MORALIS_X_API_KEY"),

		IndexerServerHost: v.GetString("INDEXER_SERVER_HOST"),

		PodtownServerHost: v.GetString("PODTOWN_SERVER_HOST"),

		RpcUrl: config.RpcUrl{
			Eth: v.GetString("ETH_RPC"),
			Ftm: v.GetString("FTM_RPC"),
			Opt: v.GetString("OPTIMISM_RPC"),
			Bsc: v.GetString("BSC_RPC"),
		},

		MarketplaceBaseUrl: config.MarketplaceBaseUrl{
			Opensea:  v.GetString("OPENSEA_BASE_URL"),
			Quixotic: v.GetString("QUIXOTIC_BASE_URL"),
			Painswap: v.GetString("PAINTSWAP_BASE_URL"),
		},
		MarketplaceApiKey: config.MarketplaceApiKey{
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

		APILayerAPIKey: v.GetString("API_LAYER_API_KEY"),
	}
}

func (v *Vault) GetString(key string) string {
	value, _ := v.data[key].(string)
	return value
}

func (v *Vault) GetBool(key string) bool {
	data, _ := v.data[key].(string)
	value, _ := strconv.ParseBool(data)
	return value
}
func (v *Vault) GetInt(key string) int {
	data, _ := v.data[key].(string)
	value, _ := strconv.Atoi(data)
	return value
}

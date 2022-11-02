# Mochi Bot API

<p align="center">
  <a href="https://getmochi.co">
    <img width="700" src="./images/mochi.png" />
  </a>
</p>

<p align="center">
  <img src="https://img.shields.io/badge/golang-1.18-blue" />
  <img src="https://img.shields.io/badge/strategy-gitflow-%23561D25" />
  <a href="https://github.com/consolelabs/mochi-api/blob/main/LICENSE">
    <img src="https://img.shields.io/badge/license-GNU-blue" />
  </a>
  <a href="https://discord.gg/DMVyp5Fr8Q">
  <img src="https://img.shields.io/github/issues-pr-raw/consolelabs/mochi-api" />
  <img src="https://img.shields.io/github/contributors/consolelabs/mochi-api" />
  <img src="https://img.shields.io/badge/we%20love-contributors-yellowgreen" />
  <a href="https://twitter.com/getmochi_bot">
  <img src="https://img.shields.io/twitter/follow/getmochi_bot?style=social" />
  </a>
</p>

Mochi Bot API is specifically designed to work with [Mochi Discord Bot](https://github.com/consolelabs/mochi-discord)

## :gear: Setup

#### Setup project

```bash
# install required cli and clone the .env
make setup

# reset and seed db
reset-db
```

### SQL Migration

```bash
# create new migration
make migrate-new name=create-new-table

# migrate up
make migration-up

# migrate down
make migration-down
```

### Start dev

```bash
make dev
```

### Run test

```bash
make test
```

## :recycle: Environment Variables

<details>
  <summary>Show table</summary>
  
  | Required | Name | Description |
  | --- | --- | --- |
  | ✅ | `BASE_URL` | Base URL that API routes will be appended on. To run in localhost, keep default env. |
  | ✅ | `PORT` | Port to run application on |
  | ✅ | `DB_HOST` | PostgreSQL IP address. All DB sample env assume you are developing locally on PostgreSQL default settings. |
  | ✅ | `DB_PORT` | PostgreSQL port |
  | ✅ | `DB_USER` | PostgreSQL username |
  | ✅ | `DB_PASS` | PostgreSQL password |
  | ✅ | `DB_NAME` | PostgreSQL database name |
  |  | `DB_SSL_MODE` | PostgreSQL SSL Mode default to false|
  |  | `TWITTER_CONSUMER_KEY` | Twitter account consumer key, all Twitter credentials are for Twitter sales bot to function, leave them blank if you don't need it |
  |  | `TWITTER_CONSUMER_SECRET` | Twitter account consumer key secret |
  |  | `TWITTER_ACCESS_TOKEN` | Twitter account access token |
  |  | `TWITTER_ACCESS_TOKEN_SECRET` | Twitter account access token secret |
  | ✅ | `CENTRALIZED_WALLET_MNEMONIC` | Mnemonic for Mochi centralized wallet, required to use any function of crypto tip bot (tipping, check balances, deposit)|
  | ✅ | `DISCORD_TOKEN` | Without this the bot cannot run. To get one, refer to [Mochi-discord](https://github.com/consolelabs/mochi-discord#readme) |
  | ✅ | `MOCHI_BOT_SECRET` | Client secret of your Discord Bot |
  |  | `JWT_SECRET` | Secret string for JWT, sample env works but you should change this |
  | ✅ | `REDIS_URL` | Redis full URL, this is for caching various data |
  |  | `MOCHI_GUILD_ID` | The guild where all Mochi's internal logs are sent to, default to the Web3 Console guild. If you're developing locally, you should set it to your test guild's id |
  |  | `LOG_CHANNEL_ID` | All logs are sent to this channel, set to your test channel if developing locally |
  |  | `MOCHI_SALE_CHANNEL_ID` | All sales bot messages are sent to this channel, set to your test channel if developing locally |
  |  | `MOCHI_ACTIVITY_CHANNEL_ID` | All guild activities messages are sent to this channel, set to your test channel if developing locally |
  |  | `MOCHI_FEEDBACK_CHANNEL_ID` | All feedbacks using `feedback` command are sent to this channel, set to your test channel if developing locally |
  |  | `OPENSEA_API_KEY` | Opensea API key, required for Opensea services. To get one, refer to [Opensea](https://docs.opensea.io/reference/request-an-api-key)|
  |  | `QUIXOTIC_API_KEY` | Quixotic Optimism API key, required for Quixotic services. To get one, refer to [Quixotic](https://quixotic.readme.io/reference/api-overview) |
  |  | `GOOGLE_CLOUD_BUCKET_NAME` | GCP Storage bucket name, we use GCP to store images with file formats that can't be rendered by Discord |
  |  | `GOOGLE_CLOUD_PROJECT_ID` | GCP project id |
  |  | `GCP_SERVICE_ACCOUNT` | GCP service account, this string is made by converting the content of GCP service account json key to `base64` |
  |  | `COVALENT_API_KEY` | Covalent API key, required for Covalent services. Toget one, refer to [Covalent](https://www.covalenthq.com/docs/api/#/0/0/USD/1) |
  |  | `INDEXER_SERVER_HOST` | Base URL of our indexer service. This is used to get NFT details. You can try our developer endpoint at or setup your own [Indexer](https://github.com/consolelabs/indexer-api) |
  |  | `PROCESSOR_SERVER_HOST` | Base URL of our indexer service. This is used to calculate user profile data. For more info, refer to Indexer |
  |  | `PODTOWN_SERVER_HOST` | Base URL of our Podtown service. This is mainly used for integrating NFT and tokens. For more info, refer to Indexer |
</details>

## :robot: API Structure

```
    Endpoints/Routes    (define endpoints, example: "GET: nft/collections/details")
            ⬇
         Handlers       (define expected response, status code, swagger)
            ⬇
         Entities       (implement all logics)
            ⬇
       ------------
        ⬇       ⬇
       Repo   Service   (database operations or calls to third party APIs, abi...)
```

## :chains: Third-party services

- CoinGecko API: used to get crypto data
- Binance API: used to get exchange history
- GCP Storage: used to store unpopular image file formats after converting them
- Marketplace API: used to query NFT collections
  - Opensea: NFT collections on Ethereum chain
  - Paintswap: NFT collections on Fantom chain
  - Quixotic: NFT collections on Optimism chain
- Covalent API: used to get historical token prices
- Indexer API: used to get NFT tokens details
- Processor API: used to get profile data
- Discord API: used to send messages directly to Discord guilds, mainly for logging
- Twitter API: used for Twitter sales bot feature

## :bulb: Conventions

### Model

Database models are defined in `pkg/model`, each table is a separated file

### Request

All API request bodies are defined in `pkg/request`, files are separated by categories

### Response

All API response bodies are defined in `pkg/response`, files are separated by categories

### Constants

Global string or number constants should be placed in `pkg/consts`, files are separated by categories

### Cronjobs

Jobs that run on a repeating schedule are defined in `pkg/jobs`

## :compass: Basic Guides

### How to: Create a new migration

Migration name should be the same as the table it's creating, if migration is not creating a new table, its name should be `alter_<table_name>_<altered_column>`

```
make migrate-new name=<migration_name>
```

If you also plan to add a new repo, add a new folder for it in `pkg/repo`, then add the new instance to `pkg/repo/repo.go` and `pkg/repo/pg/repo.go`

### How to: Create an API

1. Define your endpoints in `pkg/routes`

```go
guildGroup := v1.Group("/guilds")
{
    guildGroup.GET("", h.GetGuilds) // GET API at "<base_url>/guilds", with handler function Getguilds
}
```

2. Define API response in `pkg/handler`

#### Comments are required for swagger to generate doc

```go
// GetGuilds     godoc
// @Summary     Get guilds
// @Description Get guilds
// @Tags        Guild
// @Accept      json
// @Produce     json
// @Success     200 {object} response.GetGuildsResponse
// @Router      /guilds [get]
func (h *Handler) GetGuilds(c *gin.Context) {
	guilds, err := h.entities.GetGuilds()
	if err != nil {
		h.log.Error(err, "[handler.GetGuilds] - failed to get all guilds") // all errors should be logged
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, guilds)
}
```

3. Implement logic in `pkg/entities`

```go
func (e *Entity) GetGuilds() (*response.GetGuildsResponse, error) {
	guilds, err := e.repo.DiscordGuilds.Gets()
	if err != nil {
		return nil, err
	}

	var res response.GetGuildsResponse
	res.Data = make([]*response.GetGuildResponse, 0)
	for _, g := range guilds {
		res.Data = append(res.Data, &response.GetGuildResponse{
			ID:           g.ID,
			Name:         g.Name,
			BotScopes:    g.BotScopes,
			Alias:        g.Alias,
			LogChannelID: g.GuildConfigInviteTracker.ChannelID,
			LogChannel:   g.LogChannel,
			Active:       g.Active,
		})
	}

	return &res, nil
}
```

### How to: Create a testcase

1. Insert your database seed in `migrations/test_seed` for fixture testing
2. Use integrated Go test generator to generate a test function
3. Load up a test config with the instances needed for that specific function

```go
  db := testhelper.LoadTestDB("../../migrations/test_seed")
	repo := pg.NewRepo(db)
	cfg := config.LoadTestConfig()
	log := logger.NewLogrusLogger()
	entity := entities.New(cfg, log, repo, nil, nil, nil, nil, nil, nil, nil, nil)
	h := &Handler{
		entities: entity,
		log:      log,
	}
```

4. Add expected response in `pkg/handler/testdata`
5. Add testcases

```go
tests := []struct {
		name             string
		args             string
		wantCode         int
		wantResponsePath string
	}{
    {
			name:             "Has record guild",
			args:             "552427722551459840",
			wantCode:         http.StatusOK,
			wantResponsePath: "testdata/guilds/200-exist-discord-guild.json",
		},
		{
			name:             "400_empty_guild_id",
			wantCode:         http.StatusBadRequest,
			args:             "",
			wantResponsePath: "testdata/400-missing-guildID.json",
		}
  }
```

6. Run `make test`, make sure you have run `make migrate-test` at least once

## :pray: Credits

A big thank to all who contributed to this project!

If you'd like to contribute, please check out the [contributing guide](CONTRIBUTING.md).

[![Contributing](https://contrib.rocks/image?repo=consolelabs/mochi-api)](https://github.com/consolelabs/mochi-api/graphs/contributors)

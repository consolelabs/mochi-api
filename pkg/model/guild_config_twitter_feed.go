package model

type GuildConfigTwitterFeed struct {
	GuildID                  string `json:"guild_id"`
	TwitterConsumerKey       string `json:"twitter_consumer_key"`
	TwitterConsumerSecret    string `json:"twitter_consumer_secret"`
	TwitterAccessToken       string `json:"twitter_access_token"`
	TwitterAccessTokenSecret string `json:"twitter_access_token_secret"`
}

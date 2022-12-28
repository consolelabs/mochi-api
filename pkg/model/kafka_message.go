package model

type KafkaMessage struct {
	Platform string       `json:"platform"`
	Gitbook  GitbookClick `json:"gitbook"`
}

type GitbookClick struct {
	Command    string `json:"command"`
	Subcommand string `json:"subcommand"`
}

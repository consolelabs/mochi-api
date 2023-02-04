package request

type SnapshotEvent struct {
	ID     string `json:"id"`
	Event  string `json:"event"`
	Space  string `json:"space"`
	Expire int64  `json:"expire"`
}

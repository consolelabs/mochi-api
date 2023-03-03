package response

import "time"

type SnapshotProposalData struct {
	ID      string                `json:"id"`
	Title   string                `json:"title"`
	Body    string                `json:"body"`
	Choices []string              `json:"choices"`
	Scores  []float64             `json:"scores"`
	State   string                `json:"state"`
	Author  string                `json:"author"`
	Start   int64                 `json:"start"`
	End     int64                 `json:"end"`
	Space   SnapshotProposalSpace `json:"space"`
}

type SnapshotProposalDataResponse struct {
	Proposal *SnapshotProposalData `json:"proposal"`
}

type SnapshotSpaceDataResponse struct {
	Space *SnapshotProposalSpace `json:"space"`
}

type SnapshotProposalSpace struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CommonwealthDiscussion struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Plaintext string    `json:"plaintext"`
	Chain     string    `json:"chain"`
	Kind      string    `json:"kind"`
	Stage     string    `json:"stage"`
	CreatedAt time.Time `json:"created_at"`
}

type CommonwealthThreadResponse struct {
	Status string `json:"status"`
	Result struct {
		Threads *[]CommonwealthDiscussion `json:"threads"`
		Count   int64                     `json:"count"`
	}
}

type CommonwealthCommunity struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Website     string `json:"website"`
	IconURL     string `json:"icon_url"`
}

type CommonwealthCommunityResult struct {
	Communities []CommonwealthCommunity `json:"communities"`
	Count       int64                   `json:"count"`
}

type ListCommonwealthCommunities struct {
	Status string                       `json:"status"`
	Result *CommonwealthCommunityResult `json:"result"`
}

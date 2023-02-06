package response

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

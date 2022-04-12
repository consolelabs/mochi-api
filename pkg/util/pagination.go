package util

type Pagination struct {
	Page  int `json:"page" form:"page"`
	Size  int `json:"size" form:"size"`
	Total int `json:"total"`
}

func (p *Pagination) Standardize() {
	if p.Page < 0 {
		p.Page = 0
	}

	if p.Size <= 0 || p.Size >= 50 {
		p.Size = 24
	}
}

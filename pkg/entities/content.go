package entities

import (
	"encoding/json"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/response"
)

func (e *Entity) GetContentByType(contentType string) (*response.ProductMetadataCopy, error) {
	content, err := e.repo.Content.GetContentByType(contentType)
	if err != nil {
		e.log.Fields(logger.Fields{"type": contentType}).Errorf(err, "[entity.GetContentByType] - e.repo.Content.GetContentByType failed")
		return nil, err
	}

	var description map[string]interface{}
	err = json.Unmarshal(content.Description, &description)
	if err != nil {
		description = make(map[string]interface{})
	}

	return &response.ProductMetadataCopy{
		Id:          content.Id,
		Type:        content.Type,
		Description: description,
		CreatedAt:   content.CreatedAt,
		UpdatedAt:   content.UpdatedAt,
	}, nil
}

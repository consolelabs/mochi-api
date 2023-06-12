package entities

import (
	"time"

	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	earninfo "github.com/defipod/mochi/pkg/repo/earn_info"
	userearn "github.com/defipod/mochi/pkg/repo/user_earn"
	"github.com/defipod/mochi/pkg/request"
	"github.com/defipod/mochi/pkg/response"
)

func (e *Entity) CreateEarnInfo(req *request.CreateEarnInfoRequest) (*model.EarnInfo, error) {
	earnInfo := model.EarnInfo{
		Title:      req.Title,
		Detail:     req.Detail,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		DeadlineAt: req.DeadlineAt,
	}

	if req.PrevEarnId != nil {
		earnInfo.PrevEarnId = req.PrevEarnId
	}

	if req.DeadlineAt != nil {
		earnInfo.DeadlineAt = req.DeadlineAt
	}

	earn, err := e.repo.EarnInfo.Create(&earnInfo)
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.CreateEarnInfo] - e.repo.EarnInfo.Create failed")
		return nil, err
	}

	return earn, nil
}

func (e *Entity) GetEarnInfoList(req request.PaginationRequest) (*response.EarnInfoListResponse, error) {
	earnInfos, total, err := e.repo.EarnInfo.List(earninfo.ListQuery{
		Offset: int(req.Page * req.Size),
		Limit:  int(req.Size),
	})
	if err != nil {
		return nil, err
	}

	return &response.EarnInfoListResponse{
		Data:  earnInfos,
		Page:  int(req.Page),
		Size:  int(req.Size),
		Total: total,
	}, nil
}

func (e *Entity) CreateUserEarn(req *request.CreateUserEarnRequest) (*model.UserEarn, error) {
	userEarn, err := e.repo.UserEarn.UpsertOne(&model.UserEarn{
		UserId:     req.UserId,
		EarnId:     req.EarnId,
		Status:     req.Status,
		IsFavorite: req.IsFavorite,
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.CreateUserEarn] - e.repo.UserEarn.UpsertOne failed")
		return nil, err
	}

	return userEarn, nil
}

func (e *Entity) GetUserEarnInfoListByUserId(req request.GetUserEarnListByUserIdRequest) (*response.UserEarnListResponse, error) {
	userEarns, total, err := e.repo.UserEarn.GetByUserId(userearn.ListQuery{
		UserId: req.UserId,
		Status: req.Status,
		Limit:  int(req.Size),
		Offset: int(req.Size * req.Page),
	})
	if err != nil {
		e.log.Fields(logger.Fields{"req": req}).Errorf(err, "[entity.GetUserEarnInfoListByUserId] - e.repo.UserEarn.GetByUserId failed")
		return nil, err
	}

	return &response.UserEarnListResponse{
		Data:  userEarns,
		Page:  int(req.Page),
		Size:  int(req.Size),
		Total: total,
	}, nil

}

func (e *Entity) RemoveUserEarn(req request.DeleteUserEarnRequest) error {
	_, err := e.repo.UserEarn.Delete(&model.UserEarn{
		EarnId: req.EarnId,
		UserId: req.UserId,
	})

	return err
}

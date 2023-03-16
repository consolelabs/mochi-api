package user_token_support_request

import "github.com/defipod/mochi/pkg/model"

type Store interface {
	List(q ListQuery) ([]model.UserTokenSupportRequest, error)
	Get(id int) (*model.UserTokenSupportRequest, error)
	Create(request *model.UserTokenSupportRequest) error
	CreateWithHook(req *model.UserTokenSupportRequest, afterCreateFn func(id int) error) error
	Update(request *model.UserTokenSupportRequest) error
	UpdateWithHook(request *model.UserTokenSupportRequest, afterUpdateFn func(id int) error) error
	Delete(id int) error
}

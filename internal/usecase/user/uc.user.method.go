package user

import (
	"context"
	"fmt"

	"github.com/androsyz/inventory/consts"
	"github.com/androsyz/inventory/internal/model"
	zlog "github.com/androsyz/inventory/internal/pkg/log"
)

func (uc *Usecase) CreateUser(ctx context.Context, payload *model.CreateUserReq) (*model.UserRes, error) {
	user := &model.User{
		Name: payload.Name,
	}

	currentID, err := uc.repoUser.CreateUser(ctx, user)
	if err != nil {
		zlog.Error(ctx, nil, fmt.Sprintf(consts.ERR_CALL, "CreateUser", err))
		return nil, err
	}

	res := &model.UserRes{
		ID:   currentID,
		Name: user.Name,
	}

	return res, nil
}

func (uc *Usecase) GetUsers(ctx context.Context) (*model.UserListRes, error) {
	users, err := uc.repoUser.GetUsers(ctx)
	if err != nil {
		zlog.Error(ctx, nil, fmt.Sprintf(consts.ERR_CALL, "GetUsers", err))
		return nil, err
	}

	data := make([]*model.UserRes, 0)
	for _, u := range users {
		user := &model.UserRes{
			ID:   u.ID,
			Name: u.Name,
		}

		data = append(data, user)
	}

	res := &model.UserListRes{Users: data}
	return res, nil
}

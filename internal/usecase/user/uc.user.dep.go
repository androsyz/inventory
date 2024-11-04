package user

import (
	"context"

	"github.com/androsyz/inventory/internal/model"
)

type repoUserInterface interface {
	CreateUser(ctx context.Context, user *model.User) (int, error)
	GetUsers(ctx context.Context) ([]*model.User, error)
}

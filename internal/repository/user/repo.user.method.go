package user

import (
	"context"
	"fmt"

	"github.com/androsyz/inventory/internal/model"
	zlog "github.com/androsyz/inventory/internal/pkg/log"
)

func (repo *Repository) CreateUser(ctx context.Context, user *model.User) (int, error) {
	var currentInsertedID int

	err := repo.sqlConn.QueryRowContext(ctx, sqlInsertUser, &user.Name).Scan(&currentInsertedID)
	if err != nil {
		zlog.Error(ctx, nil, fmt.Sprintf("error when create user, got %v", err))
		return 0, err
	}

	return currentInsertedID, nil
}

func (repo *Repository) GetUsers(ctx context.Context) ([]*model.User, error) {
	var users = []*model.User{}

	rows, err := repo.sqlConn.QueryContext(ctx, sqlGetUsers)
	if err != nil {
		zlog.Error(ctx, nil, fmt.Sprintf("error when get users, got %v", err))
		return nil, err
	}

	for rows.Next() {
		temp := model.User{}
		err := rows.Scan(&temp.ID, &temp.Name)
		if err != nil {
			zlog.Error(ctx, nil, fmt.Sprintf("error when scan user, got %v", err))
		}

		users = append(users, &temp)
	}

	return users, nil
}

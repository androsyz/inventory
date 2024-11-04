package user

import (
	"bytes"
	"context"
	"errors"
	reflect "reflect"
	"testing"

	"github.com/androsyz/inventory/config"
	"github.com/androsyz/inventory/internal/model"
	zlog "github.com/androsyz/inventory/internal/pkg/log"
	gomock "github.com/golang/mock/gomock"
	"github.com/rs/zerolog"
)

func init() {
	buffer := &bytes.Buffer{}
	zl := zerolog.New(buffer).With().Timestamp().Logger()
	zlog.New(zl)
}

func TestUsecase_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockrepoUserInterface(ctrl)

	ctx := context.Background()

	type fields struct {
		cfg      *config.Config
		repoUser repoUserInterface
	}
	type args struct {
		ctx     context.Context
		payload *model.CreateUserReq
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		mockFn   func()
		wantData *model.UserRes
		wantErr  bool
	}{
		{
			"Test case successful create user",
			fields{
				repoUser: mockRepo,
			},
			args{
				ctx: ctx,
				payload: &model.CreateUserReq{
					Name: "Andro",
				},
			},
			func() {
				mockRepo.EXPECT().CreateUser(ctx, &model.User{
					Name: "Andro",
				}).Return(1, nil)
			},
			&model.UserRes{
				ID:   1,
				Name: "Andro",
			},
			false,
		},
		{
			"Test case failed create user",
			fields{
				repoUser: mockRepo,
			},
			args{
				ctx: ctx,
				payload: &model.CreateUserReq{
					Name: "Andro",
				},
			},
			func() {
				mockRepo.EXPECT().CreateUser(ctx, &model.User{
					Name: "Andro",
				}).Return(0, errors.New("error"))
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()
			uc := &Usecase{
				cfg:      tt.fields.cfg,
				repoUser: tt.fields.repoUser,
			}
			gotData, err := uc.CreateUser(ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("uc.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("uc.CreateUser() = %+v, want %+v", gotData, tt.wantData)
			}
		})
	}
}

func TestUsecase_GetUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockrepoUserInterface(ctrl)

	ctx := context.Background()

	type fields struct {
		cfg      *config.Config
		repoUser repoUserInterface
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		mockFn   func()
		wantData *model.UserListRes
		wantErr  bool
	}{
		{
			"Test case successful get users",
			fields{
				repoUser: mockRepo,
			},
			args{
				ctx: ctx,
			},
			func() {
				mockRepo.EXPECT().GetUsers(ctx).Return([]*model.User{
					{
						ID:   1,
						Name: "Andro",
					},
				}, nil)
			},
			&model.UserListRes{
				Users: []*model.UserRes{
					{
						ID:   1,
						Name: "Andro",
					},
				},
			},
			false,
		},
		{
			"Test case failed get users",
			fields{
				repoUser: mockRepo,
			},
			args{
				ctx: ctx,
			},
			func() {
				mockRepo.EXPECT().GetUsers(ctx).Return(nil, errors.New("error"))
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()
			uc := &Usecase{
				cfg:      tt.fields.cfg,
				repoUser: tt.fields.repoUser,
			}
			gotData, err := uc.GetUsers(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("uc.GetUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("uc.GetUsers() = %+v, want %+v", gotData, tt.wantData)
			}
		})
	}
}

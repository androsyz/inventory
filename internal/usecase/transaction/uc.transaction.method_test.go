package transaction

import (
	"bytes"
	"context"
	"database/sql"
	"errors"
	reflect "reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
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

func TestUsecase_CreateTransaction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepoTrx := NewMockrepoTransactionInterface(ctrl)
	mockRepoProduct := NewMockrepoProductInterface(ctrl)
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating sqlmock: %v", err)
	}
	defer mockDB.Close()

	ctx := context.Background()

	type fields struct {
		cfg             *config.Config
		sqlConn         *sql.DB
		repoTransaction repoTransactionInterface
		repoProduct     repoProductInterface
	}
	type args struct {
		ctx     context.Context
		payload *model.CreateTransactionReq
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		mockFn   func()
		wantData *model.TransactionRes
		wantErr  bool
	}{
		{
			"Test case successful create transaction",
			fields{
				repoTransaction: mockRepoTrx,
				repoProduct:     mockRepoProduct,
				sqlConn:         mockDB,
			},
			args{
				ctx: ctx,
				payload: &model.CreateTransactionReq{
					UserID:    1,
					ProductID: 1,
					Quantity:  1,
				},
			},
			func() {
				mock.ExpectBegin().WillReturnError(nil)
				mockRepoProduct.EXPECT().GetProductByID(ctx, gomock.Any(), int(1)).Return(&model.Product{
					ID:    1,
					Stock: 10,
					Price: 35000,
				}, nil).AnyTimes()
				mockRepoProduct.EXPECT().UpdateProductStock(ctx, gomock.Any(), int(9), int(1)).Return(nil).AnyTimes()
				mockRepoTrx.EXPECT().CreateTransaction(ctx, gomock.Any(), gomock.Any()).Return(1, nil).AnyTimes()
				mock.ExpectCommit().WillReturnError(nil)
			},
			&model.TransactionRes{
				ID:         1,
				UserID:     1,
				ProductID:  1,
				TotalPrice: 35000,
				Quantity:   1,
			},
			false,
		},
		{
			"Test case failed create transaction",
			fields{
				repoTransaction: mockRepoTrx,
				repoProduct:     mockRepoProduct,
				sqlConn:         mockDB,
			},
			args{
				ctx: ctx,
				payload: &model.CreateTransactionReq{
					UserID:    1,
					ProductID: 1,
					Quantity:  1,
				},
			},
			func() {
				mock.ExpectBegin().WillReturnError(errors.New("error"))
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()
			uc := &Usecase{
				cfg:             tt.fields.cfg,
				sqlConn:         tt.fields.sqlConn,
				repoProduct:     tt.fields.repoProduct,
				repoTransaction: tt.fields.repoTransaction,
			}
			gotData, err := uc.CreateTransaction(ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("uc.CreateTransaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("uc.CreateTransaction() = %+v, want %+v", gotData, tt.wantData)
			}
		})
	}
}

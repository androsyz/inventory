package supplier

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

	mockRepo := NewMockrepoSupplierInterface(ctrl)
	ctx := context.Background()

	payload := &model.CreateSupplierReq{
		Name:        "Hijab Square",
		LeadtimeMax: 7,
		LeadtimeAvg: 4,
	}

	input := &model.Supplier{
		Name:        "Hijab Square",
		LeadtimeMax: 7,
		LeadtimeAvg: 4,
	}

	type fields struct {
		cfg          *config.Config
		repoSupplier repoSupplierInterface
	}
	type args struct {
		ctx     context.Context
		payload *model.CreateSupplierReq
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		mockFn   func()
		wantData *model.SupplierRes
		wantErr  bool
	}{
		{
			"Test case successful create supplier",
			fields{
				repoSupplier: mockRepo,
			},
			args{
				ctx:     ctx,
				payload: payload,
			},
			func() {
				mockRepo.EXPECT().CreateSupplier(ctx, input).Return(1, nil)
			},
			&model.SupplierRes{
				ID:          1,
				Name:        "Hijab Square",
				LeadtimeMax: 7,
				LeadtimeAvg: 4,
			},
			false,
		},
		{
			"Test case failed create supplier",
			fields{
				repoSupplier: mockRepo,
			},
			args{
				ctx:     ctx,
				payload: payload,
			},
			func() {
				mockRepo.EXPECT().CreateSupplier(ctx, input).Return(0, errors.New("error"))
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()
			uc := &Usecase{
				cfg:          tt.fields.cfg,
				repoSupplier: tt.fields.repoSupplier,
			}
			gotData, err := uc.CreateSupplier(ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("uc.CreateSupplier() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("uc.CreateSupplier() = %+v, want %+v", gotData, tt.wantData)
			}
		})
	}
}

func TestUsecase_GetSuppliers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockrepoSupplierInterface(ctrl)

	ctx := context.Background()

	type fields struct {
		cfg          *config.Config
		repoSupplier repoSupplierInterface
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		mockFn   func()
		wantData *model.SupplierListRes
		wantErr  bool
	}{
		{
			"Test case successful get supplier",
			fields{
				repoSupplier: mockRepo,
			},
			args{
				ctx: ctx,
			},
			func() {
				mockRepo.EXPECT().GetSuppliers(ctx).Return([]*model.Supplier{
					{
						ID:          1,
						Name:        "Hijab Square",
						LeadtimeMax: 7,
						LeadtimeAvg: 4,
					},
				}, nil)
			},
			&model.SupplierListRes{
				Suppliers: []*model.SupplierRes{
					{
						ID:          1,
						Name:        "Hijab Square",
						LeadtimeMax: 7,
						LeadtimeAvg: 4,
					},
				},
			},
			false,
		},
		{
			"Test case failed get supplier",
			fields{
				repoSupplier: mockRepo,
			},
			args{
				ctx: ctx,
			},
			func() {
				mockRepo.EXPECT().GetSuppliers(ctx).Return(nil, errors.New("error"))
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()
			uc := &Usecase{
				cfg:          tt.fields.cfg,
				repoSupplier: tt.fields.repoSupplier,
			}
			gotData, err := uc.GetSuppliers(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("uc.GetSuppliers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("uc.GetSuppliers() = %+v, want %+v", gotData, tt.wantData)
			}
		})
	}
}

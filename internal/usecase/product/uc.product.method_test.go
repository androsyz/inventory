package product

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

func TestUsecase_CreateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockrepoProductInterface(ctrl)
	ctx := context.Background()

	payload := &model.CreateProductReq{
		SupplierID: 1,
		Name:       "Hijab Square Pink",
		Stock:      10,
		Price:      35000,
	}

	input := &model.Product{
		SupplierID: 1,
		Name:       "Hijab Square Pink",
		Stock:      10,
		Price:      35000,
	}

	type fields struct {
		cfg         *config.Config
		repoProduct repoProductInterface
	}
	type args struct {
		ctx     context.Context
		payload *model.CreateProductReq
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		mockFn   func()
		wantData *model.ProductRes
		wantErr  bool
	}{
		{
			"Test case successful create product",
			fields{
				repoProduct: mockRepo,
			},
			args{
				ctx:     ctx,
				payload: payload,
			},
			func() {
				mockRepo.EXPECT().CreateProduct(ctx, input).Return(1, nil)
			},
			&model.ProductRes{
				ID:         1,
				SupplierID: 1,
				Name:       "Hijab Square Pink",
				Stock:      10,
				Price:      35000,
			},
			false,
		},
		{
			"Test case failed create product",
			fields{
				repoProduct: mockRepo,
			},
			args{
				ctx:     ctx,
				payload: payload,
			},
			func() {
				mockRepo.EXPECT().CreateProduct(ctx, input).Return(0, errors.New("error"))
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()
			uc := &Usecase{
				cfg:         tt.fields.cfg,
				repoProduct: tt.fields.repoProduct,
			}
			gotData, err := uc.CreateProduct(ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("uc.CreateProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("uc.CreateProduct() = %+v, want %+v", gotData, tt.wantData)
			}
		})
	}
}
func TestUsecase_UpdateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockrepoProductInterface(ctrl)
	ctx := context.Background()

	payload := &model.UpdateProductReq{
		ID:         1,
		SupplierID: 1,
		Name:       "Hijab Square Pink",
		Stock:      10,
		Price:      36000,
	}

	input := &model.Product{
		ID:         1,
		SupplierID: 1,
		Name:       "Hijab Square Pink",
		Stock:      10,
		Price:      36000,
	}

	type fields struct {
		cfg         *config.Config
		repoProduct repoProductInterface
	}
	type args struct {
		ctx     context.Context
		payload *model.UpdateProductReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mockFn  func()
		wantErr bool
	}{
		{
			"Test case successful update product",
			fields{
				repoProduct: mockRepo,
			},
			args{
				ctx:     ctx,
				payload: payload,
			},
			func() {
				mockRepo.EXPECT().GetProductByID(ctx, nil, payload.ID).Return(input, nil)
				mockRepo.EXPECT().UpdateProduct(ctx, input).Return(nil)
			},
			false,
		},
		{
			"Test case failed update product",
			fields{
				repoProduct: mockRepo,
			},
			args{
				ctx:     ctx,
				payload: payload,
			},
			func() {
				mockRepo.EXPECT().GetProductByID(ctx, nil, payload.ID).Return(nil, errors.New("error"))
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()
			uc := &Usecase{
				cfg:         tt.fields.cfg,
				repoProduct: tt.fields.repoProduct,
			}
			err := uc.UpdateProduct(ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("uc.UpdateProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestUsecase_GetProducts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockrepoProductInterface(ctrl)

	ctx := context.Background()

	type fields struct {
		cfg         *config.Config
		repoProduct repoProductInterface
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		mockFn   func()
		wantData *model.ProductListRes
		wantErr  bool
	}{
		{
			"Test case successful get products",
			fields{
				repoProduct: mockRepo,
			},
			args{
				ctx: ctx,
			},
			func() {
				mockRepo.EXPECT().GetProducts(ctx).Return([]*model.Product{
					{
						ID:         1,
						Name:       "Hijab Square Pink",
						SupplierID: 1,
						Stock:      10,
						Price:      35000,
					},
				}, nil)
			},
			&model.ProductListRes{
				Products: []*model.ProductRes{
					{
						ID:         1,
						Name:       "Hijab Square Pink",
						SupplierID: 1,
						Stock:      10,
						Price:      35000,
					},
				},
			},
			false,
		},
		{
			"Test case failed get products",
			fields{
				repoProduct: mockRepo,
			},
			args{
				ctx: ctx,
			},
			func() {
				mockRepo.EXPECT().GetProducts(ctx).Return(nil, errors.New("error"))
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn()
			uc := &Usecase{
				cfg:         tt.fields.cfg,
				repoProduct: tt.fields.repoProduct,
			}
			gotData, err := uc.GetProducts(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("uc.GetProducts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("uc.GetProducts() = %+v, want %+v", gotData, tt.wantData)
			}
		})
	}
}

func TestUsecase_GetSafetyStock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockrepoProductInterface(ctrl)
	mockRepoSupplier := NewMockrepoSupplierInterface(ctrl)

	ctx := context.Background()

	payload := &model.SafetyStockReq{
		ProductID:  1,
		AverageReq: 5,
	}

	type fields struct {
		cfg          *config.Config
		repoProduct  repoProductInterface
		repoSupplier repoSupplierInterface
	}
	type args struct {
		ctx     context.Context
		payload *model.SafetyStockReq
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		mockFn   func()
		wantData *model.SafetyStockRes
		wantErr  bool
	}{
		{
			"Test case successful get safety stock",
			fields{
				repoProduct:  mockRepo,
				repoSupplier: mockRepoSupplier,
			},
			args{
				ctx:     ctx,
				payload: payload,
			},
			func() {
				mockRepo.EXPECT().GetProductByID(ctx, nil, payload.ProductID).Return(&model.Product{
					ID:         1,
					SupplierID: 1,
					Name:       "Hijab Square Pink",
					Stock:      10,
					Price:      35000,
				}, nil)
				mockRepoSupplier.EXPECT().GetSupplierByID(ctx, 1).Return(&model.Supplier{
					ID:          1,
					Name:        "Hijab Square",
					LeadtimeMax: 7,
					LeadtimeAvg: 4,
				}, nil)
			},
			&model.SafetyStockRes{
				ProductID:   1,
				ProductName: "Hijab Square Pink",
				Stock:       10,
				SafetyStock: 15,
			},
			false,
		},
		{
			"Test case failed get safetystock",
			fields{
				repoProduct:  mockRepo,
				repoSupplier: mockRepoSupplier,
			},
			args{
				ctx:     ctx,
				payload: payload,
			},
			func() {
				mockRepo.EXPECT().GetProductByID(ctx, nil, payload.ProductID).Return(nil, errors.New("error"))
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
				repoProduct:  tt.fields.repoProduct,
				repoSupplier: tt.fields.repoSupplier,
			}
			gotData, err := uc.GetSafetyStock(ctx, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("uc.GetSafetyStock() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("uc.GetSafetyStock() = %+v, want %+v", gotData, tt.wantData)
			}
		})
	}
}

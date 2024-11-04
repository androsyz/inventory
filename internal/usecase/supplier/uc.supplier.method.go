package supplier

import (
	"context"
	"fmt"

	"github.com/androsyz/inventory/consts"
	"github.com/androsyz/inventory/internal/model"
	zlog "github.com/androsyz/inventory/internal/pkg/log"
)

func (uc *Usecase) CreateSupplier(ctx context.Context, payload *model.CreateSupplierReq) (*model.SupplierRes, error) {
	supplier := &model.Supplier{
		Name:        payload.Name,
		LeadtimeMax: payload.LeadtimeMax,
		LeadtimeAvg: payload.LeadtimeAvg,
	}

	currentID, err := uc.repoSupplier.CreateSupplier(ctx, supplier)
	if err != nil {
		zlog.Error(ctx, nil, fmt.Sprintf(consts.ERR_CALL, "CreateSupplier", err))
		return nil, err
	}

	res := &model.SupplierRes{
		ID:          currentID,
		Name:        supplier.Name,
		LeadtimeMax: supplier.LeadtimeMax,
		LeadtimeAvg: supplier.LeadtimeAvg,
	}

	return res, nil
}

func (uc *Usecase) GetSuppliers(ctx context.Context) (*model.SupplierListRes, error) {
	suppliers, err := uc.repoSupplier.GetSuppliers(ctx)
	if err != nil {
		zlog.Error(ctx, nil, fmt.Sprintf(consts.ERR_CALL, "GetSuppliers", err))
		return nil, err
	}

	data := make([]*model.SupplierRes, 0)
	for _, s := range suppliers {
		supplier := &model.SupplierRes{
			ID:          s.ID,
			Name:        s.Name,
			LeadtimeMax: s.LeadtimeMax,
			LeadtimeAvg: s.LeadtimeAvg,
		}

		data = append(data, supplier)
	}

	res := &model.SupplierListRes{Suppliers: data}
	return res, nil
}

package model

type Supplier struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	LeadtimeMax int    `json:"leadtime_max"`
	LeadtimeAvg int    `json:"leadtime_avg"`
}

type CreateSupplierReq struct {
	Name        string `json:"name" validate:"required"`
	LeadtimeMax int    `json:"leadtime_max"`
	LeadtimeAvg int    `json:"leadtime_avg"`
}

type SupplierRes struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	LeadtimeMax int    `json:"leadtime_max"`
	LeadtimeAvg int    `json:"leadtime_avg"`
}

type SupplierListRes struct {
	Suppliers []*SupplierRes `json:"suppliers"`
}

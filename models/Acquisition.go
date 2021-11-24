package models

type GetAcquisitionBySR struct {
	SalesRetailId string `json:"sales_retail_id"`
}

type GetAcquisitionByName struct {
	Name string `json:"name"`
	RoseMerchantGroup string `json:"rose_merchant_group"`
	SalesRetailId string `json:"sales_retail_id"`
	RoseMerchantCategory string `json:"rose_merchant_category"`
	
}

type GetBusinessByCode struct {
	Code []string `json:"name"`
}

type GetAcquisitionBySalesTypeID struct {
	SalesTypeId string `json:"sales_type_id"`
}

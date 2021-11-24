package dbmodels

// ActionMerchants ..
type Acquisitions struct {
	Id                   int64  `gorm:"column:id;primary_key" json:"id"`
	Name                 string `gorm:"column:name" json:"name"`
	BusinessTypes        string `gorm:"column:business_types" json:"business_types"`
	RoseMerchantGroup    string `gorm:"column:rose_merchant_group" json:"rose_merchant_group"`
	RoseMerchantCategory string `gorm:"column:rose_merchant_category" json:"rose_merchant_category"`
	Logo                 string `gorm:"column:logo" json:"logo"`
	RegisterUsingId      string `gorm:"column:register_using_id" json:"register_using_id"`
}

// ActionMerchants ..
type AcquisitionsSR struct {
	Id            int64  `gorm:"column:id;primary_key" json:"id"`
	SalesRetailId string `gorm:"column:sales_retails" json:"sales_retails"`
}

// TableName ..
func (t *Acquisitions) TableName() string {
	return "public.acquisitions"
}

// TableName ..
func (t *AcquisitionsSR) TableName() string {
	return "public.acquisitions"
}

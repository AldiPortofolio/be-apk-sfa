package models

import "ottosfa-api-apk/database/dbmodels"

// Village ..
type Village struct {
	Village *[]dbmodels.Villages `json:"villages"`
}

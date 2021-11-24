package models

import "ottosfa-api-apk/database/dbmodels"

// District ..
type District struct {
	District *[]dbmodels.Districts `json:"districts"`
}

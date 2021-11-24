package models

import "ottosfa-api-apk/database/dbmodels"

// Province ..
type Province struct {
	Provinces *[]dbmodels.Provinces `json:"provinces"`
}

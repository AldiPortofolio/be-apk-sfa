package models

import "ottosfa-api-apk/database/dbmodels"

// City ..
type City struct {
	Cities *[]dbmodels.Cities `json:"cities"`
}

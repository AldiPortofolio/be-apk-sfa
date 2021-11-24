package fds

// LongLatMerchantResponse ...
type LongLatMerchantResponse struct {
	LongLatData     TypeLongLatData `json:"longlatData"`
	ResponCode      string          `json:"responCode"`
	DescriptionCode string          `json:"DescriptionCode"`
}

// TypeLongLatData ..
type TypeLongLatData struct {
	SFALongLat      LongLatData `json:"sfaLonglat"`
	RegisterLongLat LongLatData `json:"registerLonglat"`
	LoginLongLat    LongLatData `json:"loginLonglat"`
}

// LongLatData ..
type LongLatData struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

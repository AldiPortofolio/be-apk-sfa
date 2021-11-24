package models

// ReportHistorySummaryReq ..
type ReportHistorySummaryReq struct {
	Phone    string `json:"phone" form:"phone"`
	DateFrom string `json:"date_from" form:"date_from"`
	DateTo   string `json:"date_to" form:"date_to"`
}

// ReportHistorySummaryRes ..
type ReportHistorySummaryRes struct {
	AchievementDay  int                  `json:"achievement_day,omitempty"`
	AcquisitionData []AcquisitionSummary `json:"acquisition_data,omitempty"`
	Target          int                  `json:"target,omitempty"`
}

// AcquisitionSummary ..
type AcquisitionSummary struct {
	Created   string `json:"created"`
	Address   string `json:"address"`
	Status    string `json:"status"`
	StoreName string `json:"store_name"`
}

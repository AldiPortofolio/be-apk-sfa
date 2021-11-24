package fdsmodels

// ReportHistorySummaryRes ..
type ReportHistorySummaryRes struct {
	ResponCode      string                     `json:"responCode"`
	DescriptionCode string                     `json:"DescriptionCode"`
	AcquisitionData []ReportHistorySummaryRes1 `json:"acquisitionData"`
	AchievementDay  int                        `json:"achievementDay"`
	Target          string                     `json:"target"`
}

// ReportHistorySummaryRes1 ..
type ReportHistorySummaryRes1 struct {
	Created   string `json:"created"`
	Address   string `json:"address"`
	Status    string `json:"status"`
	StoreName string `json:"StoreName"`
}

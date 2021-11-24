package models

// CheckVersionAppsReq ..
type CheckVersionAppsReq struct {
	VersionCode int    `json:"version_code" form:"version_code"`
	Role        string `json:"role" form:"role"`
}

// CheckVersionAppsRes ..
type CheckVersionAppsRes struct {
	VersionApp  int  `json:"version_app"`
	VersionApi  int  `json:"version_api"`
	ForceUpdate bool `json:"force_update"`
}

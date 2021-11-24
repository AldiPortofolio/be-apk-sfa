package models

// CheckProfilSalesFDSReq ..
type CheckProfilSalesFDSReq struct {
	Phone string
}

// CheckProfilSalesRes ..
type CheckProfilSalesRes struct {
	ResponseCode    string                 `json:"response_code"`
	SalesName       string                 `json:"sales_name"`
	Email           string                 `json:"email"`
	DescriptionCode string                 `json:"description_code"`
	Phone           string                 `json:"phone"`
	Photo           string                 `json:"photo"`
	SessionToken    string                 `json:"session_token"`
	Status          string                 `json:"status"`
	SalesId         string                 `json:"sales_id"`
	SFAID           string                 `json:"sfa_id"`
	SumMerchant     int                    `json:"sum_merchant"`
	SumVerified     int                    `json:"sum_verified"`
	AreaAquisitions ProfileAreaAquisitions `json:"area_aquisitions"`
	Position        ProfilePosition        `json:"position"`
	Description     string                 `json:"description"`
}

// ProfilePosition ..
type ProfilePosition struct {
	ID           string `json:"id"`
	RoleID       string `json:"role_id"`
	Role         string `json:"role"`
	RegionID     string `json:"region_id"`
	Region       string `json:"region"`
	BranchID     string `json:"branch_id"`
	Branch       string `json:"branch"`
	BranchOffice string `json:"branch_office"`
	AreaID       string `json:"area_id"`
	Area         string `json:"area"`
	SubAreaID    string `json:"sub_area_id"`
	SubArea      string `json:"sub_area"`
}

// ProfileAreaAquisitions ..
type ProfileAreaAquisitions struct {
	Provinces string         `json:"province"`
	City      string         `json:"city"`
	Locations []LocationsRes `json:"locations"`
}

// LocationsRes ..
type LocationsRes struct {
	Village  string `json:"village"`
	District string `json:"district"`
}

// ListLocationSales ..
type ListLocationSales struct {
	Id        int64  `json:"id"`
	FirstName string `json:first_name"`
	LastName  string `json:last_name"`
	RoleId    int64  `json:"role_id"`
	RoleName  string `json:"role_name"`
	//RegionableID   int    `json:"regionable_id"`
	//RegionableType string `json:"regionable_type"`
	RegionId     int64  `json:"region_id"`
	RegionName   string `json:"region_name"`
	BranchId     int64  `json:"branch_id"`
	BranchName   string `json:"branch_name"`
	BranchOffice string `json:"branch_office"`
	AreaId       int64  `json:"area_id"`
	AreaName     string `json:"area_name"`
	SubAreaID    int64  `json:"sub_area_id"`
	SubAreaName  string `json:"sub_area_name"`
}

// PostionSales ..
type PostionSales struct {
	PositionId      int64   `json:"position_id"`
	RoleId        	int64   `json:"role_id"`
	RoleName 		string  `json:role_name"`
	SubAreaId  		int64   `json:sub_area_id"`
	SubAreaCode  	string  `json:sub_area_code"`
	SubAreaName  	string  `json:"sub_area_name"`
	AreaId       	int64  `json:"area_id"`
	AreaCode     	string `json:"area_code"`
	AreaName     	string `json:"area_name"`
	BranchId     	int64  `json:"branch_id"`
	BranchCode   	string `json:"branch_code"`
	BranchName   	string `json:"branch_name"`
	BranchOffice 	string `json:"branch_office"`
	RegionId     	int64  `json:"region_id"`
	RegionCode   	string `json:"region_code"`
	RegionName   	string `json:"region_name"`
	SalesTypeId     int64  `json:"sales_type_id"`
	SalesTypeName   string `json:"sales_type_name"`
}
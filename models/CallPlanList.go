package models

// CallPlanListReq ..
type CallPlanListReq struct {
	StartDate string `json:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty"`
}

// CallPlanListRes ..
type CallPlanListRes struct {
	TeamLeaderName   string         `json:"team_leader_name"`
	BranchOfficeName string         `json:"branch_office_name"`
	DateToday        string         `json:"date_today"`
	CallPlanList     []ListCallPlan `json:"call_plan_list"`
}

// ListCallPlan ..
type ListCallPlan struct {
	Date                 string `json:"date"`
	VillageName          string `json:"village_name"`
	CallPlanId           int64  `json:"call_plan_id"`
	AmountTasksCompleted int    `json:"amount_tasks_completed"`
	AmountTasks          int    `json:"amount_tasks"`
	StatusDate           int    `json:"status_date"`
}

// TeamLeaderData ..
type TeamLeaderData struct {
	TeamLeaderName   string `json:"team_leader_name"`
	BranchOfficeName string `json:"branch_office_name"`
}

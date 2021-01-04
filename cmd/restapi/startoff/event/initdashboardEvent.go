package event

import "omega/internal/core"

type temp struct {
	Total     int    `json:"total"`
	CompanyID uint64 `json:"company_id"`
}

//InitdashboardEvent will create events for dashboard procedure
func InitdashboardEvent(engine *core.Engine) {
	engine.DB.Raw(`DROP EVENT IF EXISTS sam_dashboard_update_event`).Scan(&temp{})

	dashboardEvent := `	
		CREATE EVENT IF NOT EXISTS sam_dashboard_update_event
		ON SCHEDULE EVERY 3 HOUR
		STARTS CURRENT_TIMESTAMP
		DO
		CALL refresh_dashboard_mv();`
	engine.DB.Raw(dashboardEvent).Scan(&temp{})
}

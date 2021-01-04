package event

/*
//InitreportEvent will create events for report procedure
func InitreportEvent(engine *core.Engine) {

	engine.DB.Raw(`DROP EVENT IF EXISTS sam_report_update_event`).Scan(&temp{})

	reportEvent := `
		CREATE EVENT IF NOT EXISTS sam_report_update_event
		ON SCHEDULE EVERY 3 HOUR
		STARTS CURRENT_TIMESTAMP
		DO
		CALL refresh_report_mv();`
	engine.DB.Raw(reportEvent).Scan(&temp{})
}
*/

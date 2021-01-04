package view

type tempStruct struct {
	Date      string `json:"date"`
	Total     int    `json:"total"`
	CompanyID uint64 `json:"company_id"`
}

/*

//InitViewReports will create views for the reports
func InitViewReports(engine *core.Engine) {

	//creating a materialized view for reports.
	//these tables will be used for fetching the necessary data for the daily and monthly reports based on the company

	writeQuery("daily_patients_report_mv", "daily", engine)

	//daily results
	writeQuery("daily_results_report_mv", "daily", engine)

	//daily trials
	writeQuery("daily_trials_report_mv", "daily", engine)

	//Monthly patients
	writeQuery("monthly_patients_report_mv", "monthly", engine)

	//Monthly results ..is special case as we have to consider patients too.. used in dashboard
	query := `
	CREATE TABLE  monthly_results_report_mv(
		date varchar(7) NOT NULL,
		total INT NOT NULL,
		total_patients INT NOT NULL,
		total_income DOUBLE NOT NULL,
		company_id BIGINT NOT NULL,
		INDEX company_id (company_id)
	)
	`
	dropCreateMview("monthly_results_report_mv", query, engine)

	//Monthly trials
	writeQuery("monthly_trials_report_mv", "monthly", engine)

}

func writeQuery(table string, tableType string, engine *core.Engine) {

	var query string
	if tableType == "daily" {
		query = `
		CREATE TABLE  ` + table + `(
			 date DATE NOT NULL,
			 total INT NOT NULL,
			 company_id BIGINT NOT NULL,
			 INDEX company_id (company_id)
		)
		`

	} else if tableType == "monthly" {
		query = `
		CREATE TABLE  ` + table + `(
			 date varchar(7) NOT NULL,
			 total INT NOT NULL,
			 company_id BIGINT NOT NULL,
			 INDEX company_id (company_id)
		)`
	}

	dropCreateMview(table, query, engine)

}

func dropCreateMview(table string, query string, engine *core.Engine) {
	//dropping the Materialized view if it exists
	engine.DB.Raw(`DROP TABLE IF EXISTS ` + table + `;`).Scan(&temp{})
	//(re)creating the materialized view
	engine.DB.Raw(query).Scan(&tempStruct{})
}
*/

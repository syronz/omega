package view

type temp struct {
	Total     int    `json:"total"`
	CompanyID uint64 `json:"company_id"`
}

/*
//InitDasboardViews will create necessary views for the dashboard
func InitDasboardViews(engine *core.Engine) {

	//creating a dashboard materialized view.
	//this table will be used for fetching the necessary data for the dashboard based on the company

	engine.DB.Raw(`DROP TABLE IF EXISTS dashboard_mv;`).Scan(&temp{})

	dashboardMV := `
	CREATE TABLE  dashboard_mv(
		company_id BIGINT NOT NULL
		,today_patients INT NOT NULL
		,today_results INT NOT NULL
		,today_trials INT NOT NULL
		,today_balance DOUBLE NOT NULL
		,total_trials INT NOT NULL
		,male_ctn	INT NOT NULL
		,female_ctn INT NOT NULL
		,pregnant_ctn INT NOT NULL
		,age0to12 INT NOT NULL
		,age13to17 INT NOT NULL
		,age18to24 INT NOT NULL
		,age25to34 INT NOT NULL
		,age35to44 INT NOT NULL
		,age45to54 INT NOT NULL
		,age55to64 INT NOT NULL
		,agefrom65 INT NOT NULL

		,UNIQUE INDEX company_id (company_id)
	)
	`
	engine.DB.Raw(dashboardMV).Scan(&temp{})

	//view for total patients,gender, and age based on company_id
	patients := `
	CREATE OR REPLACE VIEW dashboard_patients_report AS
	SELECT
		count(*) as total,
		count(case when DATEDIFF(now(),dob )/365.25 <13 then 1 end) as age0to12,
		count(case when DATEDIFF(now(),dob )/365.25 >=13 and DATEDIFF(now(),dob )/365.25 <18   then 1 end) as age13to17,
		count(case when DATEDIFF(now(),dob )/365.25 >=18 and DATEDIFF(now(),dob )/365.25 <25   then 1 end) as age18to24,
		count(case when DATEDIFF(now(),dob )/365.25 >=25 and DATEDIFF(now(),dob )/365.25 <35   then 1 end) as age25to34,
		count(case when DATEDIFF(now(),dob )/365.25 >=35 and DATEDIFF(now(),dob )/365.25 <44   then 1 end) as age35to44,
		count(case when DATEDIFF(now(),dob )/365.25 >=45 and DATEDIFF(now(),dob )/365.25 <55   then 1 end) as age45to54,
		count(case when DATEDIFF(now(),dob )/365.25 >=55 and DATEDIFF(now(),dob )/365.25 <65   then 1 end) as age55to64,
		count(case when DATEDIFF(now(),dob )/365.25 >=65 then 1 end) as agefrom65,
		company_id
	FROM
		sam_patients
	GROUP BY company_id
	`

	engine.DB.Raw(patients).Scan(&temp{})

	//view for total trials based on company_id
	trials := `
	CREATE OR REPLACE VIEW dashboard_trials_report AS
	SELECT
		COUNT(resultTrials.id) as total,
		resultTrials.company_id as company_id
	FROM (
		SELECT
			result.id as id, result.created_at as date, result.company_id as company_id
		FROM
			sam_results result
		INNER JOIN
	    	sam_result_trials trial
		ON
			result.id=trial.result_id and trial.created_at >=CURRENT_DATE()
	)AS resultTrials
	GROUP BY company_id;`

	engine.DB.Raw(trials).Scan(&temp{})

	//view for total resulsts based on company_id
	results := `
	CREATE OR REPLACE VIEW dashboard_results_report AS
	SELECT
		COUNT(id) as total,
		COUNT(DISTINCT patient_id) as today_patients,
		company_id
	FROM
		sam_results
	WHERE
		created_at>=CURRENT_DATE()
	GROUP BY company_id`
	engine.DB.Raw(results).Scan(&temp{})

	todaySales := `
	CREATE OR REPLACE VIEW dashboard_balance_report AS
	SELECT
	(sum(total)-sum(discount)) as today_balance,
		company_id
	FROM
		sam_results
	WHERE created_at>=CURRENT_DATE()
	GROUP BY company_id`

	engine.DB.Raw(todaySales).Scan(&temp{})

	//view for count of gender based on tests based on company_id
	trialGender := `
	CREATE OR REPLACE VIEW dashboard_trialgender_report AS
	SELECT
		count(resultTrials.id) as total_trials,
		count(case when resultTrials.gender="male" then 1 end) as male_ctn,
		count(case when resultTrials.gender="female" then 1 end) as female_ctn,
		count(case when resultTrials.gender="pregnant" then 1 end) as pregnant_ctn,
		resultTrials.company_id as company_id
	FROM (
		SELECT
			result.id as id, result.created_at as date, result.gender, result.company_id as company_id
		FROM
			sam_results result
		INNER JOIN
	    	sam_result_trials trial
		ON
			result.id=trial.result_id
	)AS resultTrials
	GROUP BY company_id;`
	engine.DB.Raw(trialGender).Scan(&temp{})

}
*/

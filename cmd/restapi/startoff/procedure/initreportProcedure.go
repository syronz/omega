package procedure

import (
	"omega/internal/core"
)

//InitReportProcedure will create report procedure
func InitReportProcedure(engine *core.Engine) {

	//we delete the existing referesh_report_procedure
	engine.DB.Raw(`DROP PROCEDURE IF EXISTS refresh_report_mv;`).Scan(&temp{})

	//Preparing the ' referesh report materialized views' procedure
	dailyReportProcedure := `	
		CREATE PROCEDURE refresh_report_mv()
			BEGIN

			TRUNCATE TABLE daily_patients_report_mv;
			TRUNCATE TABLE monthly_patients_report_mv;
			TRUNCATE TABLE daily_results_report_mv;
			TRUNCATE TABLE monthly_results_report_mv;
			TRUNCATE TABLE daily_trials_report_mv;
			TRUNCATE TABLE monthly_trials_report_mv;

			INSERT INTO daily_patients_report_mv
			SELECT
				DATE(created_at) as date,
				COUNT(id) as total,
				company_id
			FROM
				sam_patients
			GROUP BY date,company_id;
			
			INSERT INTO monthly_patients_report_mv
			SELECT
				DATE_FORMAT(created_at,'%Y-%m') as date,
				COUNT(id) as total,
				company_id
			FROM
				sam_patients
			GROUP BY date,company_id;

			INSERT INTO daily_results_report_mv
			SELECT
				DATE(created_at) as date,
				COUNT(id) as total,
				company_id
			FROM
				sam_results
			GROUP BY date,company_id;

			INSERT INTO monthly_results_report_mv
			SELECT
				DATE_FORMAT(created_at,'%Y-%m') as date,
				COUNT(id) as total,
				COUNT(DISTINCT patient_id) as total_patients,
				SUM(total-discount) as total_income,
				company_id
			FROM
				sam_results
			GROUP BY date,company_id;

			INSERT INTO daily_trials_report_mv
			SELECT
				DATE(resultTrials.date) as date,
				COUNT(resultTrials.id) as total,
				resultTrials.company_id as company_id
			FROM (
				SELECT
					result.id as id, result.updated_at as date, result.company_id as company_id
				FROM
					sam_results result
				INNER JOIN
					sam_result_trials trial
				ON
					result.id=trial.result_id
			)AS resultTrials
			GROUP BY DATE(date),company_id;


			INSERT INTO monthly_trials_report_mv
			SELECT
				DATE_FORMAT(resultTrials.date, '%Y-%m') as date,
				COUNT(resultTrials.id) as total,
				resultTrials.company_id as company_id
			FROM (
				SELECT 
					result.id as id, result.updated_at as date, result.company_id as company_id 
				FROM 
					sam_results result 
				INNER JOIN 
					sam_result_trials trial 
				ON
					result.id=trial.result_id
			) AS resultTrials
			GROUP BY DATE_FORMAT(resultTrials.date,'%Y-%m'),company_id;
			
			END;`

	//executing the dashboard procedure
	engine.DB.Raw(dailyReportProcedure).Scan(&temp{})

	//calling the procedure
	//engine.DB.Raw(`call refresh_report_mv();`).Scan(&temp{})

}

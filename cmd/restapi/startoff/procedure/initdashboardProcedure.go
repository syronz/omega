package procedure

import (
	"omega/internal/core"
)

type temp struct {
	Total     int    `json:"total"`
	CompanyID uint64 `json:"company_id"`
}

//InitDashboardProcedure will create dashboard procedure
func InitDashboardProcedure(engine *core.Engine) {

	//we delete the existing referesh_dashboard_procedure
	engine.DB.Raw(`DROP PROCEDURE IF EXISTS refresh_dashboard_mv;`).Scan(&temp{})

	//Preparing the ' referesh dashboard materialized view' procedure
	dashboardProcedure := `	
		CREATE PROCEDURE refresh_dashboard_mv()
			BEGIN

			TRUNCATE TABLE dashboard_mv;

			INSERT INTO dashboard_mv
			SELECT 
				p.company_id as company_id ,
				IFNULL(r.today_patients,0) as today_patients, 
				IFNULL(r.total,0) as today_results, 
				IFNULL(t.total,0) as today_trials,
				IFNULL(b.today_balance,0) as today_balance,
				IFNULL(tg.total_trials,0) as total_trials,
				IFNULL(tg.male_ctn,0) as male_ctn,
				IFNULL(tg.female_ctn,0) as female_ctn,
				IFNULL(tg.pregnant_ctn,0) as pregnant_ctn,
				IFNULL(p.age0to12,0) as age012, IFNULL(p.age13to17,0) as age13to17, IFNULL(p.age18to24,0) as age18to24, IFNULL(p.age25to34,0) as age25to34,
				IFNULL(p.age35to44,0) as age35to44, IFNULL(p.age45to54,0) as age45to54, IFNULL(p.age55to64,0) as age55to64, IFNULL(p.agefrom65,0) as agefrom65
			FROM 
				dashboard_patients_report p
			LEFT JOIN (
				select *
				FROM 
				dashboard_trialgender_report) tg ON p.company_id=tg.company_id
			LEFT JOIN (

				SELECT 
					*
				FROM
					dashboard_results_report
				GROUP BY company_id) r ON p.company_id=r.company_id
			LEFT JOIN (
			
			SELECT * 
			FROM 
			dashboard_balance_report) b ON r.company_id=b.company_id
			LEFT JOIN (
			select *
			FROM 
			dashboard_trials_report) t ON b.company_id=t.company_id;
			
			END;`

	//executing the dashboard procedure
	engine.DB.Raw(dashboardProcedure).Scan(&temp{})

	//calling the procedure
	//engine.DB.Raw(`call refresh_dashboard_mv();`).Scan(&temp{})
}

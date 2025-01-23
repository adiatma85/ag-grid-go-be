package pond

const (
	readPondList = `
		SELECT
		    p.id,
		    p.fk_farm_id,
		    p.fk_module_id,
		    p.latest_metric_date,
		    p.latest_sampling_date,
		    p.current_fk_cycle_id,
		    p.current_cycle_name,
		    p.current_cycle_status,
		    p.is_final_harvest,
		    p.current_cycle_severity,
		    p.current_cycle_severity_rating,
		    p.current_cycle_dop_start_date,
		    p.current_cycle_dop_end_date,
		    p.current_cycle_doc_start_date,
		    p.current_cycle_doc_end_date,
		    p.current_cycle_hatchery,
		    p.current_cycle_stocking_date,
		    p.current_cycle_stocking_netto,
		    p.current_cycle_stocking_gross,
		    p.current_cycle_stocking_actual,
		    p.current_cycle_density_netto,
		    p.current_cycle_density_gross,
		    p.current_cycle_acc_feed,
		    p.current_cycle_sampling,
		    p.current_wqi_score,
		    harvest_plan_min_shrimp_size,
		    harvest_plan_max_shrimp_size,
		    harvest_plan_day_interval,
		    harvest_plan_doc_length,
		    harvest_plan_biomass_target,
		    p.name,
		    p.code,
		    p.pond_type,
		    p.area_m2,
		    p.depth_m,
		    p.length_m,
		    p.width_m,
		    p.status,
		    p.flag,
		    p.meta,
			p.disease_description,
		    p.created_at,
		    p.created_by,
		    p.updated_at,
		    p.updated_by,
		    p.deleted_at,
		    p.deleted_by
		FROM
		    pond p`

	readPondCount = `
		SELECT
		    COUNT(*)
		FROM
		    pond p`
)

package parameter

const (
	createParameter = "INSERT INTO parameter (parameter_type, `group`, name, has_prediction, input_type, field_type, input_options, hint_id, hint_en, sub_group, time_classification, notes, priority, source, `type`, unit, optional_units, display_name_id, display_name_en, export_header_id, y_range_min, y_range_max, x_range_min, x_range_max, x_legend, threshold_type, threshold_min, threshold_max, graph_type, metrics_type, data_level, json_key, table_column_key, table_source, additional_config, api_params, api_url, flag, meta, created_at, created_by) VALUES(:parameter_type, :group, :name, :has_prediction, :input_type, :field_type, :input_options, :hint_id, :hint_en, :sub_group, :time_classification, :notes, :priority, :source, :type, :unit, :optional_units, :display_name_id, :display_name_en, :y_range_min, :y_range_max, :x_range_min, :x_range_max, :x_legend, :threshold_type, :threshold_min, :threshold_max, :graph_type, :metrics_type, :data_level, :json_key, :table_column_key, :table_source, :additional_config, :api_params, :api_url, :flag, :meta, :created_at, :created_by)"

	readParameter = `
	SELECT
		p.id,
		p.parameter_type,
		p.group,
		p.name,
		p.has_prediction,
		p.input_type,
		p.field_type,
		p.input_options,
		p.hint_id,
		p.hint_en,
		p.sub_group,
		p.time_classification,
		p.notes,
		p.priority,
		p.source,
		p.type,
		p.unit,
		p.optional_units,
		p.display_name_id,
		p.display_name_en,
		p.export_header_id,
		p.y_range_min,
		p.y_range_max,
		p.x_range_min,
		p.x_range_max,
		p.x_legend,
		p.threshold_type,
		p.alert_tracking_type,
		p.threshold_min,
		p.threshold_max,
		p.extreme_changes_lower,
		p.extreme_changes_upper,
		p.alert_threshold_min,
		p.alert_threshold_max,
		p.alert_difference,
		p.graph_type,
		p.metrics_type,
		p.data_level,
		p.json_key,
		p.table_column_key,
		p.table_source,
		p.additional_config,
		p.api_params,
		p.api_url,
		p.status,
		p.flag,
		p.meta,
		p.created_at,
		p.created_by,
		p.updated_at,
		p.updated_by,
		p.deleted_at,
		p.deleted_by
	FROM
		parameter p`

	readParameterCount = `
	SELECT
		COUNT(*)
	FROM
		parameter p`

	updateParameter = `
	UPDATE parameter`
)

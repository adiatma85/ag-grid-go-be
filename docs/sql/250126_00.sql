-- DML Inserting Slope and Intercepts for the Spectrometer Parameter
INSERT INTO `parameter`
    (`id`, `parameter_type`, `group`, `name`, `time_classification`, `unit`, `display_name_id`, `display_name_en`, `y_range_min`, `y_range_max`, `threshold_type`, `threshold_min`, `threshold_max`, `graph_type`, `metrics_type`, `data_level`, `json_key`, `table_source`, `input_type`, `field_type`, `hint_id`, `hint_en`, `priority`, `input_options`, `additional_config`, `api_params`, `api_url`)
VALUES
    -- Nitrit Intercept and Nitrit Slope
    (1036, 'water_quality', 'spechtrometer', 'nitrit_slope', '', '', 'Slope', 'Slope', 0, 0, 'none', 0, 0, 'line', 'single', 'pond', 'nitritSlope', 'pond_daily_static_metrics', 'number', 'float', '', '', 0, '', '', '', ''),
    (1037, 'water_quality', 'spechtrometer', 'nitrit_intercept', '', '', 'Intercept', 'Intercept', 0, 0, 'none', 0, 0, 'line', 'single', 'pond', 'nitritIntercept', 'pond_daily_static_metrics', 'number', 'float', '', '', 0, '', '', '', ''),

    -- Fosfat Intercept and Fosfat Slope
    (1038, 'water_quality', 'spechtrometer', 'fosfat_slope', '', '', 'Slope', 'Slope', 0, 0, 'none', 0, 0, 'line', 'single', 'pond', 'fosfatSlope', 'pond_daily_static_metrics', 'number', 'float', '', '', 0, '', '', '', ''),
    (1039, 'water_quality', 'spechtrometer', 'fosfat_intercept', '', '', 'Intercept', 'Intercept', 0, 0, 'none', 0, 0, 'line', 'single', 'pond', 'fosfatIntercept', 'pond_daily_static_metrics', 'number', 'float', '', '', 0, '', '', '', ''),

    -- TAN Intercept and TAN Slope
    (1040, 'water_quality', 'spechtrometer', 'tan_slope', '', '', 'Slope', 'Slope', 0, 0, 'none', 0, 0, 'line', 'single', 'pond', 'tanSlope', 'pond_daily_static_metrics', 'number', 'float', '', '', 0, '', '', '', ''),
    (1041, 'water_quality', 'spechtrometer', 'tan_intercept', '', '', 'Intercept', 'Intercept', 0, 0, 'none', 0, 0, 'line', 'single', 'pond', 'tanIntercept', 'pond_daily_static_metrics', 'number', 'float', '', '', 0, '', '', '', '')
    ;
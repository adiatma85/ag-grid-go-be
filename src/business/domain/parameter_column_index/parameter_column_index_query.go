package parameter_column_index

const (
	readParameterColumnIndex = `
		SELECT 
			id,
			parameter_name,
			number_format,
			is_general_column,
			is_constant,
			col_index,
			sheet_name,
			formula,
			dropdown,
			is_zero_allowed,
			status,
			flag,
			meta,
			created_at,
			created_by,
			updated_at,
			updated_by,
			deleted_at,
			deleted_by
		FROM 
			parameter_column_index
	`

	countParameterColumnIndex = `
		SELECT 
			COUNT(*)
		FROM
			parameter_column_index
	`

	insertParameterColumnIndex = `
		INSERT INTO parameter_column_index
		(
			parameter_name,
		 	number_format,
		 	is_general_column,
		 	is_constant,
		 	col_index,
		 	sheet_name,
		 	formula,
		 	dropdown,
		 	created_at,
		 	created_by
		)
		VALUES
		(
			:parameter_name,
		 	:number_format,
		 	:is_general_column,
		 	:is_constant,
		 	:col_index,
		 	:sheet_name,
		 	:formula,
		 	:dropdown,
		 	:created_at,
		 	:created_by
		)
	`

	updateParameterColumnIndex = `
		UPDATE
			parameter_column_index
	`
)

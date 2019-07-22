LEFT JOIN (
	SELECT 
		identifier,
		group_concat(name||"="||value,"&") AS param,is_edit
	FROM 
		"request_data_schemas"
	GROUP BY identifier ,is_edit
	ORDER BY id ASC
) AS reqd ON 
	request_info_schemas.identifier = reqd.identifier AND request_info_schemas.is_edit = reqd.is_edit
LEFT JOIN(
	SELECT 
		req.identifier as identifier, method , path, proto, host, header, param
	FROM "request_info_schemas" as req
	LEFT JOIN (
		SELECT
			identifier,
			'{"data":{'||group_concat('"'||name||'":"'||value||'"',",")||"}}" AS param
		FROM
			"request_data_schemas"
		GROUP BY identifier
		ORDER BY id ASC
	) AS reqd ON
		req.identifier = reqd.identifier
	LEFT JOIN (
		SELECT
			identifier,
			'{"header":{'||group_concat('"'||name||'": "'||value||'"',",")||"}}" AS header
		FROM
			"request_header_schemas"
		GROUP BY identifier
		ORDER BY id ASC
	) AS reqh ON
		req.identifier = reqh.identifier
	WHERE is_edit = 0
) AS reqNoEdit ON
	history_schemas.identifier = reqNoEdit.identifier
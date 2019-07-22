LEFT JOIN(
	SELECT res.identifier,body,header
	FROM "response_info_schemas" as res
	LEFT JOIN (
		SELECT
			identifier,
			'{"header":{'||group_concat('"'||name||'":"'||value||'"',",")||"}}" AS header
		FROM
			"response_header_schemas"
		GROUP BY identifier
		ORDER BY id ASC
	) AS resh ON
		res.identifier = resh.identifier
	LEFT JOIN (
		SELECT
			identifier, body
		FROM
			"response_body_schemas"
		GROUP BY identifier
		ORDER BY id ASC
	) AS resb ON
		res.identifier = resb.identifier
) AS resA ON
history_schemas.identifier = resA.identifier
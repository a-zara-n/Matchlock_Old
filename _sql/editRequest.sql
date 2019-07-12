LEFT JOIN(
	SELECT 
		req.identifier as identifier, method, path, proto, host, header, param
	FROM "requests" as req
	LEFT JOIN (
		SELECT
			identifier,
			'{"data":{'||group_concat('"'||name||'":"'||value||'"',",")||"}}" AS param
		FROM
			"request_data"
		GROUP BY identifier
		ORDER BY id ASC
	) AS reqd ON
		req.identifier = reqd.identifier
	LEFT JOIN (
		SELECT
			identifier,
			'{"header":{'||group_concat('"'||name||'":"'||value||'"',",")||"}}" AS header
		FROM
			"request_headers"
		GROUP BY identifier
		ORDER BY id ASC
	) AS reqh ON
		req.identifier = reqh.identifier
	WHERE is_edit = 1
) AS reqEdit ON
	histories.identifier = reqEdit.identifier
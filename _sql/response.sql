LEFT JOIN(
	SELECT res.identifier,body,header
	FROM "responses" as res
	LEFT JOIN (
		SELECT
			identifier,
			'{"header":{'||group_concat('"'||name||'":"'||value||'"',",")||"}}" AS header
		FROM
			"response_headers"
		GROUP BY identifier
		ORDER BY id ASC
	) AS resh ON
		res.identifier = resh.identifier
	LEFT JOIN (
		SELECT
			identifier, body
		FROM
			"response_bodies"
		GROUP BY identifier
		ORDER BY id ASC
	) AS resb ON
		res.identifier = resb.identifier
) AS resA ON
histories.identifier = resA.identifier
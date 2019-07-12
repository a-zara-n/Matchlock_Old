LEFT JOIN (
	SELECT 
		identifier,
		group_concat(name||"="||value,"&") AS param 
	FROM 
		"request_data"
	GROUP BY identifier 
	ORDER BY id ASC
) AS reqd ON 
	requests.identifier = reqd.identifier
package httpserver

import (
	"encoding/json"
	"net/http"
)

var historyCount int
var requestDataTable = `
LEFT JOIN (
	SELECT 
		identifier,
		group_concat(name||"="||value,"&") AS param 
	FROM 
		"request_data"
	GROUP BY identifier 
	ORDER BY id ASC
) AS reqd ON 
	requests.identifier = reqd.identifier`

var httpData = map[string]string{
	"SELECT": `
distinct histories.identifier as identifier,
reqNoEdit.method    AS request_method,
reqNoEdit.path      AS request_path,
reqNoEdit.proto     AS request_proto,
reqNoEdit.host      AS request_host,
reqNoEdit.header    AS request_headers,
reqNoEdit.param     AS request_param,
reqEdit.method      AS request_edit_method,
reqEdit.path        AS request_edit_path,
reqEdit.proto       AS request_edit_proto,
reqEdit.host        AS request_edit_host,
reqEdit.header      AS request_edit_headers,
reqEdit.param       AS request_edit_param,
resA.header 		AS response_headers,
body`,
	"NoEditReq": `
LEFT JOIN(
	SELECT 
		req.identifier as identifier, method , path, proto, host, header, param
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
			'{"header":{'||group_concat('"'||name||'": "'||value||'"',",")||"}}" AS header
		FROM
			"request_headers"
		GROUP BY identifier
		ORDER BY id ASC
	) AS reqh ON
		req.identifier = reqh.identifier
	WHERE is_edit = 0
) AS reqNoEdit ON
	histories.identifier = reqNoEdit.identifier
`,
	"EditReq": `
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
`,
	"Response": `
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
`,
}

func getHistory(i int) []historyData {
	historys := []historyData{}
	hisdb := db.OpenDatabase()
	hisdb.Table("requests").
		Select(" id,requests.identifier as identifier,method,host,path,url,param").
		Joins(requestDataTable).
		Where("id >= ?", i).
		Find(&historys)
	return historys
}
func GetHistryAll(w http.ResponseWriter, r *http.Request) {
	historyCount = 1
	historys := getHistory(historyCount)
	res, err := json.Marshal(APIresponse{Data: historys})
	historyCount = len(historys)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func GetRequest(identifier string, w http.ResponseWriter, r *http.Request) {
	ret := httpdata{}
	hisdb := db.OpenDatabase()
	hisdb.Table("histories").
		Select(httpData["SELECT"]).
		Joins(httpData["NoEditReq"]).
		Joins(httpData["EditReq"]).
		Joins(httpData["Response"]).
		Where("histories.identifier = ?", identifier).
		Find(&ret)
	res, err := json.Marshal(APIresponse{Data: ret})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

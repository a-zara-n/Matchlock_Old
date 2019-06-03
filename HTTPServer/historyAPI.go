package httpserver

import (
	"encoding/json"
	"net/http"
)

func GetHistry(w http.ResponseWriter, r *http.Request) {
	historys := []historyData{}
	hisdb := db.OpenDatabase()
	hisdb.Table("requests").
		Select(" id,requests.identifier as identifier,method,host,path,url,param").
		Joins(`
		LEFT JOIN (
			SELECT 
				identifier,
				group_concat(name||"="||value,"&") AS param 
			FROM 
				"request_data"
			GROUP BY identifier 
			ORDER BY id ASC
		) AS reqd ON 
			requests.identifier = reqd.identifier`).
		Find(&historys)
	res, err := json.Marshal(APIresponse{Data: historys})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

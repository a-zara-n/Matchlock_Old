package aggregate

//HistoryDataDefinitionByJSON はWSで利用されるJSONの定義
type HistoryDataDefinitionByJSON struct {
	ID         int64  `json:"ID"`
	Identifier string `json:"Identifier"`
	Method     string `json:"Method"`
	Host       string `json:"Host"`
	Path       string `json:"Path"`
	URL        string `json:"URL"`
	Param      string `json:"Param"`
}

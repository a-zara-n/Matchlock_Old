package repository

//HistoryRepository はhttp historyの情報を取得するためのインターフェースです
type HistoryRepository interface {
	Count() int
	Fetch(i int) []HTTPHistoryDefinitionJSON
	Insert(Identifier string, IsEdit bool)
	Update(Identifier string, IsEdit bool)
}

//HTTPHistoryDefinitionJSON はAPIで利用されるJSONのschemaを定義しています
type HTTPHistoryDefinitionJSON struct {
	ID         int64  `json:"ID"`
	Identifier string `json:"Identifier"`
	Method     string `json:"Method"`
	Host       string `json:"Host"`
	Path       string `json:"Path"`
	URL        string `json:"URL"`
	Param      string `json:"Param"`
}

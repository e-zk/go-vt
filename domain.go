package vt

// domain response model
type DomainResp struct {
	Categories   map[string]string `json:"categories"`
	CreationDate int               `json:"creation_date"`
	Jarm         string            `json:"jarm"`
	LastAnalysis int               `json:"last_analysis_date"`
	LastAnalysis map[string]struct {
		Category  string `json:"category"`
		Engine    string `json:"engine_name"`
		EngineVer string `json:"engine_version"`
		Method    string `json:"method"`
		Result    string `json:"result"`
	} `json:"last_analysis_results"`
	LastAnalysisStats map[string]int `json:"last_analysis_stats"`
	// ...
}

func Domain(domain string) {

}

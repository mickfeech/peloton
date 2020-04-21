package models

type WorkOutDetails struct {
	Duration         int `json:"duration"`
	AverageSummaries []struct {
		DisplayName string `json:"display_name"`
		DisplayUnit string `json:"display_unit"`
		Slug        string `json:"slug"`
		Value       int    `json:"value"`
	} `json:"average_summaries"`
	Summaries []struct {
		DisplayName string `json:"display_name"`
		DisplayUnit string `json:"display_unit"`
		Slug        string `json:"slug"`
		Value       int    `json:"value"`
	} `json:"summaries"`
	Metrics []struct {
		AverageValue        int    `json:"average_value"`
		DisplayName         string `json:"display_name"`
		DisplayUnit         string `json:"display_unit"`
		MaxValue            int    `json:"max_value"`
		MissingDataDuration int    `json:"missing_data_duration"`
		Slug                string `json:"slug"`
		Values              []int  `json:"values"`
		Zones               []struct {
			DisplayName string `json:"display_name"`
			Duration    int    `json:"duration"`
			MaxValue    int    `json:"max_value"`
			MinValue    int    `json:"min_value"`
			Range       string `json:"range"`
			Slug        string `json:"slug"`
		} `json:"zones"`
	} `json:"metrics"`
	TargetPerformanceMetrics struct {
		CadenceTimeInRange    int `json:"cadence_time_in_range"`
		ResistanceTimeInRange int `json:"resistance_time_in_range"`
		TargetGraphMetrics    []struct {
			Average   int `json:"average"`
			GraphData struct {
				Average []int `json:"average"`
				Lower   []int `json:"lower"`
				Upper   []int `json:"upper"`
			} `json:"graph_data"`
			Max  int    `json:"max"`
			Min  int    `json:"min"`
			Type string `json:"type"`
		} `json:"target_graph_metrics"`
	} `json:"target_performance_metrics"`
}

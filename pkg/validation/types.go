package pkg

// +k8s:deepcopy-gen=true
type ValidationResponse struct {
	Result  string          `json:"result,omitempty"`
	Reason  string          `json:"reason,omitempty"`
	Message string          `json:"message,omitempty"`
	Details ResponseDetails `json:"details,omitempty"`
}

// +k8s:deepcopy-gen=true
type ResponseDetails struct {
	TotalMismatch int `json:"totalMismatch,omitempty"`

	IdsMissingFromElasticsearch      []string `json:"idsMissingFromElasticsearch,omitempty"`
	IdsMissingFromElasticsearchCount int      `json:"idsMissingFromElasticsearchCount,omitempty"`

	IdsOnlyInElasticsearch      []string `json:"idsOnlyInElasticsearch,omitempty"`
	IdsOnlyInElasticsearchCount int      `json:"idsOnlyInElasticsearchCount,omitempty"`

	IdsWithMismatchContent []string `json:"idsWithMismatchContent,omitempty"`

	MismatchContentDetails []MismatchContentDetails `json:"mismatchContentDetails,omitempty"`
}

// +k8s:deepcopy-gen=true
type MismatchContentDetails struct {
	ID                   string `json:"id,omitempty"`
	ElasticsearchContent string `json:"elasticsearchContent,omitempty"`
	DatabaseContent      string `json:"databaseContent,omitempty"`
}

package pkg

// +k8s:deepcopy-gen=true
type ValidationResponse struct {
	Result  ValidationResult `json:"result,omitempty"`
	Reason  string           `json:"reason,omitempty"`
	Message string           `json:"message,omitempty"`
	Details ResponseDetails  `json:"details,omitempty"`
}

// ValidationResult from the validation run.
// +enum
type ValidationResult string

// These are the possible validation results.
const (
	ValidationValid     ValidationResult = "valid"
	ValidationInvalid   ValidationResult = "invalid"
	ValidationNew       ValidationResult = "new"
	ValidationUndefined ValidationResult = ""
)

// +k8s:deepcopy-gen=true
type ResponseDetails struct {
	Counts  CountDetails   `json:"counts,omitempty"`
	IDs     IdsDetails     `json:"ids,omitempty"`
	Content ContentDetails `json:"content,omitempty"`
}

// +k8s:deepcopy-gen=true
type CountDetails struct {
	InconsistencyAbsolute      int     `json:"InconsistencyAbsolute"`
	InconsistencyRatio         float64 `json:"inconsistencyRatio"`
	RecordCountInElasticsearch int     `json:"recordCountInElasticsearch"`
	RecordCountInDatabase      int     `json:"recordCountInDatabase"`
}

// +k8s:deepcopy-gen=true
type IdsDetails struct {
	InconsistencyAbsolute            int      `json:"InconsistencyAbsolute"`
	InconsistencyRatio               float64  `json:"inconsistencyRatio"`
	AmountValidated                  int      `json:"amountValidated"`
	IdsMissingFromElasticsearch      []string `json:"idsMissingFromElasticsearch,omitempty"`
	IdsMissingFromElasticsearchCount int      `json:"idsMissingFromElasticsearchCount,omitempty"`
	IdsOnlyInElasticsearch           []string `json:"idsOnlyInElasticsearch,omitempty"`
	IdsOnlyInElasticsearchCount      int      `json:"idsOnlyInElasticsearchCount,omitempty"`
}

// +k8s:deepcopy-gen=true
type ContentDetails struct {
	InconsistencyAbsolute  int               `json:"InconsistencyAbsolute"`
	InconsistencyRatio     float64           `json:"inconsistencyRatio"`
	AmountValidated        int               `json:"amountValidated"`
	IdsWithMismatchContent []string          `json:"idsWithMismatchContent,omitempty"`
	MismatchContentDetails MismatchedRecords `json:"mismatchContentDetails,omitempty"`
}

// +k8s:deepcopy-gen=true
type ContentDiff struct {
	Diffs      []string `json:"diffs"`
	ESDocument string   `json:"esDocument"`
	DBRecord   string   `json:"dbRecord"`
}

func (i *ContentDiff) AddDiff(diff string) {
	i.Diffs = append(i.Diffs, diff)
}

// +k8s:deepcopy-gen=true
type MismatchedRecords map[string]*ContentDiff //map of record id to diff details

// +k8s:deepcopy-gen=true
type DatasourceDBConnection struct {
	SslMode  string `json:"sslMode,omitempty"`
	Name     string `json:"name,omitempty"`
	Hostname string `json:"hostname,omitempty"`
	Password string `json:"password,omitempty"`
	Username string `json:"username,omitempty"`
	Table    string `json:"table,omitempty"`
	Port     string `json:"port,omitempty"`
}

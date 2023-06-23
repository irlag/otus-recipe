package elastic_index

//go:generate easyjson

//easyjson:json
type Mapping struct {
	Settings map[string]interface{} `json:"settings"`
	Mappings map[string]interface{} `json:"mappings"`
}

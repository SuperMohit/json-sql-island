package resources

// Query schema in order to map the json query to the query generator
type QuerySchema struct {
	Select struct {
		Columns []string `json:"columns"`
	} `json:"select"`
	From struct {
		Tables []string `json:"tables"`
	} `json:"from"`
	Where []struct {
		Operator   string      `json:"operator"`
		Fieldname  string      `json:"fieldName"`
		Fieldvalue interface{} `json:"fieldValue"`
	} `json:"where,omitempty"`
	Join []struct {
		Type  string `json:"type"`
		Table string `json:"table"`
		On    string `json:"on"`
	} `json:"join,omitempty"`
	Group   []string `json:"group"`
	Orderby struct {
		Columns []struct {
			Name string `json:"name"`
			Desc bool   `json:"desc,omitempty"`
		} `json:"columns"`
	} `json:"orderBy,omitempty"`
}

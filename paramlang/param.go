package paramlang

type ParamValue struct {
	// Values is the user provided value verbatim.
	Value string `json:"value"`

	// Type indicates how 'Value' should be parsed.
	// Type stackItemType `json:"param_type"`
}

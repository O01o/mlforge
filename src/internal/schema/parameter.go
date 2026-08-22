package sc

type CreateParameterRequest struct {
	ParameterName  string `json:"name"`
	ParameterValue string `json:"value"`
}

type Parameter struct {
	ParameterID    uint64 `json:"id"`
	ParameterName  string `json:"name"`
	ParameterValue string `json:"value"`
}

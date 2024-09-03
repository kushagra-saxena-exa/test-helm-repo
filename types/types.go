package types

type Response struct {
	Output `json:"output"`
}

type Output struct {
	Parameters []Param `json:"parameters"`
}

type Param struct {
	SubCode string `json:"subcode"`
}

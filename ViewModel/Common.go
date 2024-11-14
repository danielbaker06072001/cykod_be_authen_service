package ViewModel

type CommonSearchResponse struct {
	Result       string `json:"result"`
	ErrorMessage string `json:"errormessage"`
	PageIndex    int64  `json:"pageindex"`
	PageSize     int64  `json:"pagesize"`
}

type CommonUpdateResponse struct {
	Result       string `json:"result"`
	IsExisted    bool   `json:"isexisted"`
	ErrorMessage string `json:"errormessage"`
	Guid         string `json:"guid"`
}

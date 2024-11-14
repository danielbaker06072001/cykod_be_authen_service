package DTO

type SearchParam struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

type AddedParam map[string]string

type AddedParamInterface map[string]interface{}

type GetPageParam struct {
	PageIndex    interface{}    `json:"pageindex"`
	PageSize     interface{}    `json:"pagesize"`
	SearchParams *[]SearchParam `json:"searchparams"`
}

type GetSearchParam struct {
	SearchParams *[]SearchParam `json:"searchparams"`
}

type ListAddedParamInterface []AddedParamInterface

type ListAddedParam []AddedParam

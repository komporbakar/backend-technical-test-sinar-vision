package response

type WebResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type WebResponseWithLimitAndOffset struct {
	Status     string      `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Limit      int         `json:"limit"`
	Offset     int         `json:"offset"`
	TotalCount int64       `json:"total_count"`
}

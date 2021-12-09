package web_response

type response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func WebResponse(code int, status string, data interface{}) response {
	return response{
		Code:   code,
		Status: status,
		Data:   data,
	}
}
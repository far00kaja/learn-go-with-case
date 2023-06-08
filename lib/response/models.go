package response

type ResponseData struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseBadRequest struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}

type ResponseNoData struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}

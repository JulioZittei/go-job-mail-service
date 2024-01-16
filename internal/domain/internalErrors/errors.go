package internalerrors

type ErrorsParam struct {
	Param string `json:"param"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Status string `json:"status"`
	Code int `json:"code"`
	Title string `json:"title"`
	Detail string `json:"detail"`
	Instance string `json:"instance"`
	InvalidParams []ErrorsParam `json:"invalid_params"`
}
package response

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

func NewSuccess(data interface{}) *Response {
	return &Response{
		Success: true,
		Data:    data,
	}
}

func NewError(err string) *Response {
	return &Response{
		Success: false,
		Data:    err,
	}
}

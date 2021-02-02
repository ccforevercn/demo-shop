package service

type Response struct {
	Data interface{}  `json:"data"`
	Message string `json:"message"`
	Code int16 `json:"code"`
}

// 成功返回
func (response *Response) GetSuccess(data interface{}, message string) Response {
	res := Response{Data: data, Message: message, Code: 200}
	return  res
}

// 失败返回
func (response *Response) GetError(data interface{}, message string) Response {
	res := Response{Data: data, Message: message, Code: 400}
	return  res
}

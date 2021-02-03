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
func (response *Response) GetError(message string) Response {
	res := Response{Message: message, Code: 400}
	return  res
}

// 指定状态返回
func (response *Response) GetStatus(data interface{}, message string, code int16) Response {
	res := Response{Data: data, Message: message, Code: code}
	return  res
}

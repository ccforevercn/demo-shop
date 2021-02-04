package response

var param map[string]interface{}

// 成功返回
func GetSuccess(data interface{}, message string) map[string]interface{} {
	param["message"] = message
	param["code"] = 200
	param["data"] = data
	return param
}

// 成功返回(没有参数)
func GetSuccessMessage(message string) map[string]interface{} {
	param["message"] = message
	param["code"] = 200
	return param
}

// 失败返回
func GetError(message string) map[string]interface{} {
	param["message"] = message
	param["code"] = 400
	return param
}

// 指定状态返回
func GetStatus(data interface{}, message string, code int16) map[string]interface{} {
	param["message"] = message
	param["code"] = code
	param["data"] = data
	return param
}


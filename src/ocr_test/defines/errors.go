package defines

const (
	RES_SUCCESS                      = 0

	RES_ERR_PARAMETERS_WRONG         = 1001
	RES_ERR_IMAGE_MISSED             = 1002
	RES_ERR_IMAGE_TYPE_NOT_SUPPORTED = 1003
	RES_ERR_IMAGE_URL_WRONG          = 1004
	RES_ERR_BASE64_WRONG             = 1005

	RES_ERR_DB_WRONG                 = 9006
	RES_ERR_UNKNOW                   = 9999
)

type ErrorInfo struct{
	Msg string
	StatusCode int
}

var ErrorInfoMap = map[int]ErrorInfo{
	RES_ERR_PARAMETERS_WRONG:         ErrorInfo{"参数错误", 417},
	RES_ERR_IMAGE_MISSED:             ErrorInfo{"图片缺失", 417},
	RES_ERR_IMAGE_TYPE_NOT_SUPPORTED: ErrorInfo{"图片类型不支持", 417},
	RES_ERR_IMAGE_URL_WRONG:          ErrorInfo{"图片下载URL不支持", 417},
	RES_ERR_BASE64_WRONG:             ErrorInfo{"图片BASE64错误", 417},
	RES_ERR_DB_WRONG:                 ErrorInfo{"数据库错误", 500},
	RES_ERR_UNKNOW:                   ErrorInfo{"未知错误", 500},
}

func GetErrorInfo(code int) ErrorInfo {
	return ErrorInfoMap[code]
}

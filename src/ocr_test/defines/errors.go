package defines

const (
	RES_SUCCESS                      = 0
	RES_ERR_PARAMETERS_WRONG         = 1001
	RES_ERR_IMAGE_MISSED             = 1002
	RES_ERR_IMAGE_TYPE_NOT_SUPPORTED = 1003
	RES_ERR_IMAGE_URL_WRONG          = 1004
	RES_ERR_BASE64_WRONG             = 1005
	RES_ERR_DB_WRONG                 = 1006
	RES_ERR_UNKNOW                   = 9999
)

var Error_map = map[int]string{
	RES_ERR_PARAMETERS_WRONG:         "Wrong Parameters.",
	RES_ERR_IMAGE_MISSED:             "图片缺失",
	RES_ERR_IMAGE_TYPE_NOT_SUPPORTED: "图片类型不支持",
	RES_ERR_IMAGE_URL_WRONG:          "图片下载URL不支持",
	RES_ERR_BASE64_WRONG:             "图片BASE64错误",
	RES_ERR_DB_WRONG:                 "数据库错误",
	RES_ERR_UNKNOW:                   "未知错误",
}

func Translate(code int) string {
	return Error_map[code]
}

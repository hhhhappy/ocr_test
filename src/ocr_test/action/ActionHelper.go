package action

import (
	"github.com/gin-gonic/gin"
	"ocr_test/defines"
	"ocr_test/manager/logManager"
	"strconv"
)

func MakeErrorResponce(context *gin.Context, code int) {
	errorInfo := defines.GetErrorInfo(code)

	context.JSON(errorInfo.StatusCode, gin.H{
		"msg": errorInfo.Msg,
		"res": code,
	})
}

func TraitImageUploadError(context *gin.Context, err error) {
	code, errConvert := strconv.Atoi(err.Error())
	if errConvert != nil {
		//Logger
		logManager.LogError(err.Error(), true)
		code = defines.RES_ERR_UNKNOW
	}

	MakeErrorResponce(context, code)
}
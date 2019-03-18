package image

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ocr_test/action"
	"ocr_test/defines"
	"ocr_test/form/imageForm"
	"ocr_test/manager/logManager"
	"ocr_test/model/tDetect"
)

func ListResultAction(context *gin.Context) {
	//Form
	var form imageForm.ListResultForm
	err := context.ShouldBind(&form)

	//Check the validity of form
	if err != nil {
		action.MakeErrorResponce(context, defines.RES_ERR_PARAMETERS_WRONG)
		return
	}

	result, err := tDetect.FindAll()
	if err != nil {
		logManager.LogError(err.Error(), false)
		action.MakeErrorResponce(context, defines.RES_ERR_DB_WRONG)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"res":     defines.RES_SUCCESS,
		"detect_result": result,
	})
}

package image

import (
	"github.com/gin-gonic/gin"
	"github.com/otiai10/gosseract"
	"net/http"
	"ocr_test/action"
	"ocr_test/defines"
	"ocr_test/form/imageForm"
	"ocr_test/manager/configManager"
	"ocr_test/manager/fileManager"
	"ocr_test/manager/logManager"
	"ocr_test/model/tDetect"
	"time"
	"unicode"
)

func DetectAction(context *gin.Context) {
	//Form
	var form imageForm.DetectForm
	err := context.ShouldBind(&form)

	//Check the validity of form
	if err != nil {
		action.MakeErrorResponce(context, defines.RES_ERR_PARAMETERS_WRONG)
		return
	}

	//Check if there is an uploaded image
	image, _ := context.FormFile("image_file");
	var fileFullPath string
	var fileName string

	dir := configManager.GetConf().FilePathRoot + defines.ImageFilePath + "/"
	if image != nil {
		//Http upload directly
		fileFullPath, fileName, err = fileManager.HttpUploadImage(dir, image, context)
	} else if form.ImageBase64 != "" {
		//Base64 upload
		fileFullPath, fileName, err = fileManager.Base64UploadImage(dir, form.ImageBase64)
	} else if form.ImageUrl != "" {
		//Url upload
		fileFullPath, fileName, err = fileManager.UrlUploadImage(dir, form.ImageUrl)
	} else {
		//Image wasn't uploaded
		action.MakeErrorResponce(context, defines.RES_ERR_IMAGE_MISSED)
		return
	}

	//Error occurred during the upload
	if err != nil {
		action.TraitImageUploadError(context, err)
		return
	}

	// if the file path is empty
	if fileFullPath == "" || fileName == "" {
		action.MakeErrorResponce(context, defines.RES_ERR_UNKNOW)
		return
	}

	content := []string{}

	// detector initial
	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage(fileFullPath)
	text, _ := client.Text()
	for _, char:= range text  {
		if (unicode.IsLetter(char)) {	// if it's letter
			content = append(content, string(char))
		}
	}

	detect := tDetect.TDetect{
		Content:    content,
		DetectTime: time.Now().Format("Mon Jan 2 15:04:05"),
		ImageFile:  fileFullPath,
	}

	_, err = detect.Insert()
	if err != nil {
		logManager.LogError(err.Error(), false)
		action.MakeErrorResponce(context, defines.RES_ERR_DB_WRONG)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"res":     defines.RES_SUCCESS,
		"content": content,
	})
}

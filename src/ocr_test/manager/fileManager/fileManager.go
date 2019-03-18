package fileManager

import (
	"bytes"
	"encoding/base64"
	"errors"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	. "ocr_test/defines"
	"ocr_test/utility"
	"os"
	"strconv"
	"strings"
)

/*Upload Image by HTTP*/
func HttpUploadImage(path string, file *multipart.FileHeader, c *gin.Context) (string, string, error) {

	//verify file's extension
	res := strings.Split(file.Filename, ".")
	extension := strings.ToLower(res[len(res)-1])
	if _, ok := IMAGE_TYPE_SUPPORTED[extension]; !ok {
		return "", "", errors.New(strconv.Itoa(RES_ERR_IMAGE_TYPE_NOT_SUPPORTED))
	}

	//reset the filename by using UUID
	fileName := utility.GetUUID() + "." + extension
	fileFullPath := path + fileName

	//check if the directory exists
	res2, err := CreateDir(path)
	if res2 == false {
		return "", "", err
	}

	//save file
	err = c.SaveUploadedFile(file, fileFullPath)

	if err != nil {
		return "", "", err
	}

	//set the file as 777
	err = os.Chmod(fileFullPath, 0777)
	if err != nil {
		return "", "", err
	}

	return fileFullPath, fileName, nil
}

/*Upload Image by Url*/
func UrlUploadImage(path string, url string) (fileFullPath string, fileName string, err error) {
	extension := "jpg"

	response, err := http.Get(url)
	defer response.Body.Close()

	if err != nil {
		return "", "", errors.New(strconv.Itoa(RES_ERR_IMAGE_URL_WRONG))
	}

	//reset the filename by using UUID
	fileName = utility.GetUUID() + "." + extension
	fileFullPath = path + fileName

	//check if the directory exists
	res2, err := CreateDir(path)
	if res2 == false {
		return "", "", err
	}

	//store the file
	out, err := os.Create(fileFullPath)
	defer out.Close()

	pix, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", "", err
	}

	_, err = io.Copy(out, bytes.NewReader(pix))
	if err != nil {
		return "", "", err
	}

	//set the file as 777
	err = os.Chmod(fileFullPath, 0777)
	if err != nil {
		return "", "", err
	}

	return fileFullPath, fileName, nil
}

/*Upload Image by BASE64*/
func Base64UploadImage(path string, base64String string) (fileFullPath string, fileName string, err error) {
	extension := "jpg"

	//reset the filename by using UUID
	fileName = utility.GetUUID() + "." + extension
	fileFullPath = path + fileName

	//check if the directory exists
	res2, err := CreateDir(path)
	if res2 == false {
		return "", "", err
	}

	//base64 encoding
	stream, err := base64.StdEncoding.DecodeString(base64String)

	if err != nil {
		return "", "", errors.New(strconv.Itoa(RES_ERR_BASE64_WRONG))
	}

	//save file
	err = ioutil.WriteFile(fileFullPath, stream, 0664)
	if err != nil {
		return "", "", err
	}

	//set the file as 777
	err = os.Chmod(fileFullPath, 0777)
	if err != nil {
		return "", "", err
	}

	return fileFullPath, fileName, nil
}

func CreateDir(dir string) (bool, error) {
	_, err := os.Stat(dir)

	if err == nil {
		//directory exists
		return true, nil
	}

	//not exists
	err2 := os.MkdirAll(dir, 0777)
	if err2 != nil {
		return false, err2
	}

	return true, nil
}

func DeleteFIle(path string) error {
	return os.Remove(path)
}

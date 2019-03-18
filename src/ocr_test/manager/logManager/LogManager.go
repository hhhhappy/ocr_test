package logManager

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"ocr_test/defines"
	"ocr_test/manager/configManager"
	"ocr_test/manager/fileManager"
	"ocr_test/manager/mongoManager"
	"runtime"
	"time"
	"os"
	"log"
	"strings"
	"strconv"
)

/* MongoDB log writer */
type MongoWriter struct {
	session *mgo.Session
}

/* File writer */
var fileWriter *os.File

func (mw *MongoWriter) Write(p []byte) (n int, err error) {
	c := mw.session.DB("db_ocr_test").C("t_system_errors_log")
	err = c.Insert(bson.M{
		"datetime": time.Now().Format("2006-01-02 15:04:05"),
		"message": string(p),
		"channel":"Golang",
	})
	if err != nil {
		return
	}
	return len(p), nil
}

/* save log to file*/
func logToFile(msg string)  {
	if fileWriter == nil {
		//get configs
		conf := configManager.GetConf()

		logFilePath := conf.FilePathRoot + defines.LogFilePath

		//check if the directory exists
		_, err := fileManager.CreateDir(logFilePath)
		if err != nil {
			panic(err)
		}
		fileWriter, err = os.OpenFile(logFilePath + "error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
	}

	log.SetOutput(fileWriter)
	log.SetPrefix("[Error]")
	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)

	_, callerFile, line, _ := runtime.Caller(2)
	log.Println(fmt.Sprintf("%s, IN FILE: %s, ON LINE: %d", msg, callerFile, line))
}

/* save log to file*/
func logToMongoDb(msg string)  {
	mw := &MongoWriter{mongoManager.GetSession()}
	log.SetOutput(mw)
	log.SetPrefix("[Error]")
	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)

	_, callerFile, line, _ := runtime.Caller(2)
	log.Println(fmt.Sprintf("%s, IN FILE: %s, ON LINE: %d", msg, callerFile, line))
}

/* Log error */

func LogError(msg string, ifLogToMongo bool)  {
	logToFile(msg)
	if ifLogToMongo {
		logToMongoDb(msg)
	}
}

/*  */

func MakeText(msg string, txtToInsert []string) string {
	for key, txt := range txtToInsert {
		msg = strings.Replace(msg, "::" + strconv.Itoa(key), txt, 1)
	}
	return msg
}

func GetLogFileWriter() *os.File {
	if fileWriter == nil {
		//get configs
		conf := configManager.GetConf()

		logFilePath := conf.FilePathRoot + defines.LogFilePath

		//check if the directory exists
		_, err := fileManager.CreateDir(logFilePath)
		if err != nil {
			panic(err)
		}
		fileWriter, err = os.OpenFile(logFilePath + "error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
	}
	return fileWriter
}
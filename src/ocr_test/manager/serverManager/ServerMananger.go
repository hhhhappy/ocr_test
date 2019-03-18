package serverManager

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"ocr_test/action/image"
	"time"
)

var router *gin.Engine
var server *http.Server

func Initial()  {
	router = gin.Default()

	router.Use(func(context *gin.Context) {
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Methods", "GET, POST")
		context.Header("Access-Control-Allow-Headers", "x-requested-with,content-type")
	})
}

/* load the Api Url */
func LoadRouter()  {
	//define image routes
	imageGroup :=router.Group("/image")
	{
		imageGroup.POST("/detect", image.DetectAction)
		imageGroup.POST("/list_result", image.ListResultAction)
	}
}

func CreateServer(port string) {
	server = &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
}

func StartToListen() error {
	return server.ListenAndServe()
}

func ServerShutdown()  {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		panic("can't stop gracefully the server!")
	}
}

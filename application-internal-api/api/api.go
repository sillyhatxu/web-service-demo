package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sillyhatxu/gin-utils/v2"
	"github.com/sillyhatxu/gin-utils/v2/entity"
	"github.com/sillyhatxu/gin-utils/v2/gincodes"
	"github.com/sillyhatxu/gin-utils/v2/response"
	"github.com/sillyhatxu/web-service-demo/common/handler"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strconv"
	"time"
)

func InitialAPI(port int) {
	router := SetupRouter()
	err := router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		logrus.Errorf("server down. %v", err)
		panic(err)
	}
}

func SetupRouter() *gin.Engine {
	router, err := ginutils.SetupRouter()
	if err != nil {
		panic(err)
	}
	router.Use(handler.TimeoutMiddleware(30 * time.Second))
	demoGroup := router.Group("/demos")
	{
		demoGroup.GET("/slow/:second", slow)
		demoGroup.GET("/fast", fast)
	}
	envGroup := router.Group("/envs")
	{
		envGroup.GET("/get/:key", get)
	}
	internalGroup := router.Group("/internal-api")
	{
		internalGroup.GET("/get", getInternalDetail)
	}
	return router
}

func fast(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, response.Success(entity.Data(nil)))
	return
}

func slow(ctx *gin.Context) {
	secondParam := ctx.Param("second")
	i, err := strconv.Atoi(secondParam)
	if err != nil {
		ctx.JSON(http.StatusOK, response.Errorf(gincodes.ParamsValidateError, err))
		return
	}
	time.Sleep(time.Duration(i) * time.Second)
	ctx.JSON(http.StatusOK, response.Success(entity.Data(nil)))
	return
}

func get(ctx *gin.Context) {
	keyParam := ctx.Param("key")
	ctx.JSON(http.StatusOK, response.Success(entity.Data(os.Getenv(keyParam))))
	return
}

func getInternalDetail(ctx *gin.Context) {
	mockDetail := os.Getenv("MOCK_DETAIL")
	if mockDetail == "" {
		mockDetail = fmt.Sprintf(`{"value":"%v","time":"%v"}`, "hello world", time.Now().Format(time.RFC3339))
	}
	var data map[string]interface{}
	err := json.Unmarshal([]byte(mockDetail), &data)
	if err != nil {
		ctx.JSON(http.StatusOK, response.Errorf(gincodes.ParamsValidateError, err))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(entity.Data(data)))
	return
}

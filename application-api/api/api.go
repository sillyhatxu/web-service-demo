package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sillyhatxu/gin-utils/v2"
	"github.com/sillyhatxu/gin-utils/v2/entity"
	"github.com/sillyhatxu/gin-utils/v2/gincodes"
	"github.com/sillyhatxu/gin-utils/v2/response"
	client "github.com/sillyhatxu/http-client"
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
		envGroup.GET("/get/:key", getEnv)
	}
	internalGroup := router.Group("/internal")
	{
		internalGroup.GET("/get", getInternal)
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

func getEnv(ctx *gin.Context) {
	keyParam := ctx.Param("key")
	ctx.JSON(http.StatusOK, response.Success(entity.Data(os.Getenv(keyParam))))
	return
}

func getInternal(ctx *gin.Context) {
	host := os.Getenv("INTERNAL_HOST")
	if host == "" {
		ctx.JSON(http.StatusOK, response.NewError(gincodes.ServerError, "unset INTERNAL_HOST"))
		return
	}
	httpClient := client.NewHttpClient(host)
	httpResponse, err := httpClient.DoGet("/internal-api/get")
	if err != nil {
		ctx.JSON(http.StatusOK, response.Errorf(gincodes.ServerError, err))
		return
	}
	var res DemoResponse
	err = httpClient.AnalysisBody(httpResponse, &res)
	if err != nil {
		ctx.JSON(http.StatusOK, response.Errorf(gincodes.ServerError, err))
		return
	}
	if res.Code != gincodes.OK {
		ctx.JSON(http.StatusOK, response.NewError(res.Code, res.Msg))
		return
	}
	ctx.JSON(http.StatusOK, response.Success(entity.Data(res.Data)))
	return
}

type DemoResponse struct {
	Code  string       `json:"code"`
	Data  *interface{} `json:"data"`
	Msg   string       `json:"message"`
	Extra *interface{} `json:"extra"`
}

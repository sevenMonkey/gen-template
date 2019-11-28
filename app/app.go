package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"{{projectName}}/pkg/e"
)


type Resp struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func RespJson(c *gin.Context, errCode int, data interface{} )  {
	if data == nil{
		data = make(map[string]interface{})
	}
	c.JSON(http.StatusOK, Resp{Code:errCode, Data:data, Msg:e.GetMsg(errCode)})
}

func RespRawJson(c *gin.Context, data []byte)  {
	c.Data(http.StatusOK, "Content-Type:application/json", data)
}

func NoRoute(c *gin.Context) {
	RespJson(c, e.ERROR_INVALID_REQUEST, nil)
}
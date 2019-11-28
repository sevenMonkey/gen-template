package pkg

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"{{projectName}}/app"
	"{{projectName}}/app/logging"
	"{{projectName}}/pkg/e"
)

func ParseRequest(c *gin.Context, request interface{}) error {
	err := c.ShouldBindWith(request, binding.JSON)
	if err != nil {
		app.RespJson(c, e.ERROR_INVLIAD_PARA, nil)
		logging.Info("ParseRequest Result", request)
		logging.Info("ParseRequest Error", err.Error())
		return err
	}
	return nil
}

func SuccessResponse(c *gin.Context, response interface{}) {
	c.JSON(http.StatusOK, response)
}

func CheckErr(c *gin.Context, err error) {
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
}

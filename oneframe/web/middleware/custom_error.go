package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yumosx/oneframe/lv_log"
	"github.com/yumosx/oneframe/utils/lv_err"
	"github.com/yumosx/oneframe/web/lv_dto"
)

func RecoverError(c *gin.Context) {
	defer func() {
		err := recover()
		if err != nil {
			switch errTypeObj := err.(type) {
			case string:
				if strings.HasPrefix(errTypeObj, "{") {
					c.Header("Content-Type", "application/json; charset=utf-8")
					c.String(http.StatusOK, errTypeObj)
					c.Abort()
				} else {
					c.JSON(http.StatusOK, gin.H{"code": 500, "msg": errTypeObj})
					c.Abort()
				}
			case lv_dto.Resp:
				c.AbortWithStatusJSON(http.StatusOK, errTypeObj)
			case error:
				if gin.IsDebugging() {
					lv_err.PrintStackTrace(errTypeObj)
				}
				lv_log.Error(c, "CustomError: ", errTypeObj)
				c.JSON(http.StatusOK, gin.H{"code": 500, "msg": errTypeObj.Error()})
				c.Abort()
			default:
				lv_log.Error(c, "CustomError: ", errTypeObj)
				c.JSON(http.StatusOK, gin.H{"code": 500, "msg": "未知错误!"})
				c.Abort()
			}
		} else {
			lv_log.Info(c, "request over")
		}
	}()
	c.Next()
}

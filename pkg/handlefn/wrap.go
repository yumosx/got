package handlefn

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yumosx/got/pkg/code"
)

// S 表示 shadow 的意思, 表示需要去处理返回参数
func S[Resp any](fn func(ctx *gin.Context) (code.Result[Resp], error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res, err := fn(ctx)
		if err != nil { // 用户产生的问题, 我们希望用户对返回值做处理
			ctx.PureJSON(http.StatusInternalServerError, res)
			return
		}
		ctx.JSON(http.StatusOK, res)
	}
}

// P 表示有参数的意思
func P[Req any, Resp any](fn func(ctx *gin.Context, req Req) (code.Result[Resp], error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req Req
		if err := ctx.Bind(&req); err != nil {
			ctx.PureJSON(http.StatusInternalServerError, code.InError())
			return
		}
		res, err := fn(ctx, req)
		if err != nil {
			ctx.PureJSON(http.StatusInternalServerError, res)
			return
		}
		ctx.JSON(http.StatusOK, res)
	}
}

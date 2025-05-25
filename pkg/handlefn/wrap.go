package handlefn

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yumosx/got/pkg/code"
	"github.com/yumosx/got/pkg/errx"
)

// S 表示 shadow 的意思, 不需要去处理参数
func S[T any](fn func(ctx *gin.Context) errx.Option[code.Result[T]]) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res := fn(ctx)
		if res.NoNil() { //用户产生的问题, 我们希望用户对返回值做处理
			ctx.PureJSON(http.StatusInternalServerError, res.Val)
			return
		}
		ctx.JSON(http.StatusOK, res.Val)
	}
}

// R 表示有参数的意思, 需要去处理参数
func R[T any](fn func(ctx *gin.Context, req T) errx.Option[code.Result[T]]) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req T
		if err := ctx.Bind(&req); err != nil {
			ctx.PureJSON(http.StatusInternalServerError, code.InError())
			return
		}
		res := fn(ctx, req)
		if res.NoNil() {
			ctx.PureJSON(http.StatusInternalServerError, res.Val)
			return
		}
		ctx.JSON(http.StatusOK, res.Val)
	}
}

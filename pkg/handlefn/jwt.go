package handlefn

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type JwtMiddlewareBuilder struct {
	paths  []string
	secret string
	claims jwt.Claims
}

func NewJwtMiddlewareBuilder(claims jwt.Claims, secret string) *JwtMiddlewareBuilder {
	return &JwtMiddlewareBuilder{claims: claims, secret: secret}
}

func (j *JwtMiddlewareBuilder) IgnorePaths(path string) *JwtMiddlewareBuilder {
	j.paths = append(j.paths, path)
	return j
}

func (j *JwtMiddlewareBuilder) Build(renew func(ctx *gin.Context, claims jwt.Claims) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, path := range j.paths {
			if path == c.Request.URL.Path {
				return
			}
		}

		// 1. 获取对应的token
		tokenHeader := c.GetHeader("Authorization")
		if tokenHeader == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		sets := strings.SplitN(tokenHeader, " ", 2)

		if len(sets) != 2 && sets[0] != "Bearer" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		tokenStr := sets[1]

		// 2. 解析对应的token
		token, err := jwt.ParseWithClaims(tokenStr, j.claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// 3. 刷新的逻辑交个客户端
		err = renew(c, j.claims)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}

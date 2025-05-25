package handlefn

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type CorsMiddlewareBuilder struct {
	config cors.Config
}

func NewCorsMiddlewareBuilder() *CorsMiddlewareBuilder {
	return &CorsMiddlewareBuilder{config: cors.Config{}}
}

func (c *CorsMiddlewareBuilder) IgnorePath() {
}

func (c *CorsMiddlewareBuilder) Build() gin.HandlerFunc {
	return cors.New(c.config)
}

package middleware

import (
	"github.com/opentracing/opentracing-go/ext"
	"blog/global"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/gin-gonic/gin"
	"context"
)

func Tracing() func(c *gin.Context) {
	return func (c *gin.Context)  {
		var newCtx context.Context
		var span opentracing.Span
		spanCtx, err := opentracing.GlobalTracer().Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(c.Request.Header),
		)
		if err != nil {
			span, newCtx = opentracing.StartSpanFromContextWithTracer(
				c.Request.Context(),
				global.Tracer,
				c.Request.URL.Path,
			)
		} else {
			span, newCtx = opentracing.StartSpanFromContextWithTracer(
				c.Request.Context(),
				global.Tracer,
				c.Request.URL.Path,
				opentracing.ChildOf(spanCtx),
				opentracing.Tag{Key: string(ext.Component), Value: "http"},
			)
		}
		defer span.Finish()
		c.Request = c.Request.WithContext(newCtx)
		c.Next()
	}
}
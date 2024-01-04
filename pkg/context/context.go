package context

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"net/url"
)

func InitGinContext() *gin.Context {
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request = &http.Request{
		URL:    &url.URL{},
		Header: map[string][]string{},
	}
	return ctx
}

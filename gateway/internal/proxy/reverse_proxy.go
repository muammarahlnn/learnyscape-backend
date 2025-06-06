package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/muammarahlnn/learnyscape-backend/gateway/internal/log"
	"github.com/muammarahlnn/learnyscape-backend/pkg/constant"
	ginutil "github.com/muammarahlnn/learnyscape-backend/pkg/util/gin"
)

func NewReverseProxy(target string) gin.HandlerFunc {
	url, err := url.Parse(target)
	if err != nil {
		log.Logger.Errorf("failed parsing url: %v", url)
	}
	proxy := httputil.NewSingleHostReverseProxy(url)

	return func(ctx *gin.Context) {
		defer func() {
			if err, ok := recover().(error); ok && err != nil {
				ctx.Error(err)
				ctx.Abort()
			}
		}()

		params := map[string]any{
			"path":   ctx.Request.URL.Path,
			"target": target,
		}
		log.Logger.WithFields(params).Info("proxying request")

		proxy.Director = func(req *http.Request) {
			req.URL.Scheme = url.Scheme
			req.URL.Host = url.Host
			req.URL.Path = ctx.Param("path")
			req.Header = ctx.Request.Header

			if userID := ginutil.GetUserID(ctx); userID != 0 {
				req.Header.Set(constant.X_USER_ID, strconv.FormatInt(userID, 10))
			}

			if role := ginutil.GetUserRole(ctx); role != "" {
				req.Header.Set(constant.X_ROLE, role)
			}
		}

		proxy.ServeHTTP(ctx.Writer, ctx.Request)
	}
}

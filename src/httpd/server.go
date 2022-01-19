package httpd

import (
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"net/http"
)

const HTTP_PORT = "8000"

//HttpServer
func HttpServer(logger *zap.Logger) {
	s := &fasthttp.Server{
		Handler: routeLoader().Handler,
	}
	logger.Info("HTTP server start")
	err := s.ListenAndServe(":" + HTTP_PORT)
	if err != nil {
		logger.Error("Error in ListenAndServe:" + err.Error())
	}
}

//RouteLoader
func routeLoader() *fasthttprouter.Router {
	router := fasthttprouter.New()
	router.GET("/health", handleHealth)
	router.GET("/ready", handleReady)
	return router
}

func handleHealth(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusOK)
	fmt.Fprint(ctx, "{\"status\": \"OK\"}")
}

func handleReady(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusOK)
	fmt.Fprint(ctx, http.StatusText(http.StatusOK))
}

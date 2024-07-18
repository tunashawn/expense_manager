package response

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

type HttpResponse struct {
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
}

type JWTToken struct {
	Token string `json:"token"`
}

func (r HttpResponse) Response(code int, data interface{}, ctx *gin.Context) {
	ctx.JSON(code, data)
}

func (r HttpResponse) Success(data interface{}, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, HttpResponse{
		Data:    data,
		Message: "ok",
	})
}

func (r HttpResponse) Unauthorized(err error, ctx *gin.Context) {
	slog.Error("Unauthorized Request", "err", err.Error())
	ctx.JSON(http.StatusUnauthorized, HttpResponse{Message: err.Error()})
}

func (r HttpResponse) BadRequest(err error, ctx *gin.Context) {
	slog.Error("Bad Request", "err", err.Error())
	ctx.JSON(http.StatusBadRequest, HttpResponse{Message: err.Error()})
}

func (r HttpResponse) InternalServerError(err error, ctx *gin.Context) {
	slog.Error("Internal Server Error", "err", err.Error())
	ctx.JSON(http.StatusInternalServerError, HttpResponse{Message: err.Error()})
}

package httpapi

import (
	"example.com/m/v2/errs"
	"github.com/gin-gonic/gin"
)

type EndPoint interface {
	GetType() int
	GetSubPath() string
	GetHandle() gin.HandlerFunc
}

type BaseJson struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func BaseError(err *errs.CodeError) *BaseJson {
	return &BaseJson{err.Code(), err.Error()}
}

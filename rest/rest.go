package rest

import (
	"github.com/gin-gonic/gin"
)

func RunAPI(address string) error {
	// Gin 엔진
	r := gin.Default()
	r.GET("/", IndexHome)

	// run server
	return r.Run(address)
}

/*
	원본코드
	h := HandlerInterface
	r.GET("/", h.IndexHome)
	핸들러 사용시 에러 발생. 추후 해결 예정
*/

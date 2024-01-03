package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"lancer/constant"
	"lancer/response"
	"net/http"
)

//log request body in context

func RequestLog() gin.HandlerFunc {
	return func(c *gin.Context) {

		bodyBytes, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusAccepted, response.ResultData{Code: http.StatusAccepted, Msg: err.Error(), Data: nil})
			c.Abort()
			return
		}

		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		c.Set(constant.RequestBody, string(bodyBytes))

		c.Next()
	}
}

func RequestLogContext(c *gin.Context) {
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusAccepted, response.ResultData{Code: http.StatusAccepted, Msg: err.Error(), Data: nil})
		c.Abort()
		return
	}

	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	c.Set(constant.RequestBody, string(bodyBytes))
}

package controller

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"lancer/constant"
	request "lancer/request/ping"
	response2 "lancer/response"
	response "lancer/response/ping"
	"net/http"
	"time"
)

type PingController struct {
}

func NewPingController() *PingController {
	return &PingController{}
}

func (controller *PingController) Ping(c *gin.Context) {

	response := make(map[string]interface{})

	response["code"] = http.StatusOK
	response["msg"] = "ping"
	response["time"] = time.Now().Format(time.DateTime)
	response["data"] = ""
	c.JSON(http.StatusOK, response)
	return
}

func (controller *PingController) Empty(c *gin.Context) {
	req := new(request.PingEmptyRequest)
	ret := new(response.PingEmptyResponse)

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.Set(constant.ResponseBody, response2.ResultData{Code: http.StatusBadRequest, Msg: err.Error(), Data: nil})
		return
	}

	if len(body) != 0 || len(string(body)) != 0 {
		// Put the requested weight back in for subsequent processing
		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
		err := c.ShouldBindJSON(&req)
		if err != nil {
			c.Set(constant.ResponseBody, response2.ResultData{Code: http.StatusBadRequest, Msg: err.Error(), Data: nil})
			return
		}
	}

	//TODO

	c.Set(constant.ResponseBody, ret)
	return

}

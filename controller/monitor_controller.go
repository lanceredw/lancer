package controller

import (
	"github.com/gin-gonic/gin"
	"lancer/constant"
	"lancer/plugin/translate"
	request "lancer/request/monitor"
	response2 "lancer/response"
	response "lancer/response/monitor"

	"lancer/service"
	"net/http"
)

type MonitorController struct {
}

func NewMonitorController() *MonitorController {
	return &MonitorController{}
}

func (controller *MonitorController) Goroutines(c *gin.Context) {
	req := new(request.MonitorGoroutinesRequest)
	ret := new(response.MonitorGoroutinesResponse)

	err := c.ShouldBindQuery(&req)
	if err != nil {
		c.Set(constant.ResponseBody, response2.ResultData{Code: http.StatusAccepted, Msg: translate.Translate(err), Data: nil})
		return
	}

	srv := service.NewMonitorService()
	ret, err = srv.Goroutines(req)
	if err != nil {
		c.Set(constant.ResponseBody, response2.ResultData{Code: http.StatusAccepted, Msg: err.Error(), Data: nil})
		return
	}
	c.Set(constant.ResponseBody, ret)
	return
}

func (controller *MonitorController) FileStats(c *gin.Context) {
	req := new(request.MonitorFileStatsRequest)
	ret := new(response.MonitorFileStatsResponse)

	err := c.ShouldBindQuery(&req)
	if err != nil {
		c.Set(constant.ResponseBody, response2.ResultData{Code: http.StatusAccepted, Msg: translate.Translate(err), Data: nil})
		return
	}

	srv := service.NewMonitorService()
	ret, err = srv.FileStats(req)
	if err != nil {
		c.Set(constant.ResponseBody, response2.ResultData{Code: http.StatusAccepted, Msg: err.Error(), Data: nil})
		return
	}
	c.Set(constant.ResponseBody, ret)
	return
}

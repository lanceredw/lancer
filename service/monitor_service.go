package service

import (
	"lancer/common/method"
	"lancer/global"
	request "lancer/request/monitor"
	response "lancer/response/monitor"
	"net/http"
	"runtime"
)

type MonitorService struct {
}

func NewMonitorService() *MonitorService {
	return &MonitorService{}
}

func (service *MonitorService) Goroutines(req *request.MonitorGoroutinesRequest) (ret *response.MonitorGoroutinesResponse, err error) {
	ret = new(response.MonitorGoroutinesResponse)

	num := runtime.NumGoroutine()

	ret.Code = http.StatusOK
	ret.Msg = global.TranslateMessage["FindSuccess"]
	ret.Data.Num = num
	return
}

func (service *MonitorService) FileStats(req *request.MonitorFileStatsRequest) (ret *response.MonitorFileStatsResponse, err error) {
	ret = new(response.MonitorFileStatsResponse)

	//run in linux
	stats, err := method.ParseFileFDStats("/proc/sys/fs/file-nr")
	if err != nil {
		return nil, err
	}

	ret.Code = http.StatusOK
	ret.Msg = global.TranslateMessage["FindSuccess"]
	ret.Data = stats

	return
}

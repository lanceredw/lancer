package scheduled

import (
	"github.com/robfig/cron/v3"
	"lancer/global"
)

var (
	Cron *cron.Cron
)

func NewCron() {
	Cron = cron.New(cron.WithSeconds())

	//run the timer
	Cron.Start()

	//load the tasks
	LoadTasks()
}

func LoadTasks() {

	CronJobs()

	//TODO OnceJobs()
}

func CronJobs() {

	//每天晚上进行日志滚动
	logSpec := "10 0 0 * * *"
	logEntryID, err := Cron.AddFunc(logSpec, func() {
		StartRotate()
	})
	if err != nil {
		global.Logger.Error("添加定时任务错误:", err)
	} else {
		global.Logger.Info("添加定时任务日志滚动成功:", logEntryID)
	}
}

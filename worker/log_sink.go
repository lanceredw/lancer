package worker

import (
	"lancer/global"
	"lancer/model"
)

func InitLogSink() {

	GLogSink = &LogSink{
		LogChan: make(chan *model.LancerActionLog, 1000),
	}

	go GLogSink.writeLoop()

	return
}

//write Log

func (logSink *LogSink) writeLoop() {
	var (
		logOne *model.LancerActionLog
	)

	for {
		select {
		case logOne = <-logSink.LogChan:
			logSink.SaveLog(logOne)

		}
	}
}

func (logSink *LogSink) SaveLog(log *model.LancerActionLog) {
	err := global.DB.Model(&model.LancerActionLog{}).Create(&log).Error
	if err != nil {
		global.Logger.Error("logging failed")
	}
}

func (logSink *LogSink) Send(log *model.LancerActionLog) {
	logSink.LogChan <- log
}

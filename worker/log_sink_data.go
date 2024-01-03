package worker

import "lancer/model"

var (
	GLogSink *LogSink
)

type LogSink struct {
	LogChan chan *model.LancerActionLog
}

package response

type MonitorFileStatsResponse struct {
	Code int               `json:"code"`
	Msg  string            `json:"msg"`
	Data map[string]string `json:"data"` //jwt
}

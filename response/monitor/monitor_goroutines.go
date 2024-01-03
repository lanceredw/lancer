package response

type MonitorGoroutinesResponse struct {
	Code int        `json:"code"`
	Msg  string     `json:"msg"`
	Data Goroutines `json:"data"` //jwt
}

type Goroutines struct {
	Num  int    `json:"num"`
	Info string `json:"info"`
}

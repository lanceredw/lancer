package data

type RedisData struct {
	Url      string `json:"url"`
	Password string `json:"password"`
	Db       int    `json:"db"`
}

package data

type JwtData struct {
	Secret string `json:"secret"`
	Expire int64  `json:"expire"`
}

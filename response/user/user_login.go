package response

type UserLoginResponse struct {
	Code int     `json:"code"`
	Msg  string  `json:"msg"`
	Data JwtData `json:"data"` //jwt
}

type JwtData struct {
	Token string `json:"token"`
}

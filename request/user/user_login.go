package request

type UserLoginRequest struct {
	UserName string `json:"user_name" form:"user_name" binding:"required"` //user name
	Password string `json:"password" form:"password" binding:"required"`   //password
}

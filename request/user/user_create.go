package request

type UserCreateRequest struct {
	UserName string `json:"user_name" form:"user_name" binding:"required" label:"user name"`
	Password string `json:"password" form:"password" binding:"required" label:"password"`
}

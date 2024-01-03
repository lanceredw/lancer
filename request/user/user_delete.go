package request

type UserDeleteRequest struct {
	Id int64 `json:"id,string" form:"id,string" binding:"required"` //user id
}

package request

type RbacMenuDeleteRequest struct {
	Id int64 `json:"id,string" form:"id,string" binding:"required"` //menu id
}

package request

type RbacRoleDeleteRequest struct {
	Id int64 `json:"id,string" form:"id,string" binding:"required"` //role id
}

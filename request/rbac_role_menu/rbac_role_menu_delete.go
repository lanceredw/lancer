package request

type RbacRoleMenuDeleteRequest struct {
	Id int64 `json:"id,string" form:"id,string" binding:"required"` //role_menu id
}

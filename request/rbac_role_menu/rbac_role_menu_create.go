package request

type RbacRoleMenuCreateRequest struct {
	RoleId int64 `json:"role_id,string" form:"role_id,string" binding:"required"` //role id
	MenuId int64 `json:"menu_id,string" form:"menu_id,string" binding:"required"` //menu id
}

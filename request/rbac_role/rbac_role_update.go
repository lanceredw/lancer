package request

type RbacRoleUpdateRequest struct {
	Id       int64  `json:"id,string" form:"id,string" binding:"required" label:"id"`        //role id
	RoleName string `json:"role_name" form:"role_name" binding:"required" label:"role name"` //role name
}

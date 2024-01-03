package request

type RbacRoleCreateRequest struct {
	RoleName string `json:"role_name" form:"role_name" binding:"required" label:"role name"` //role name
}

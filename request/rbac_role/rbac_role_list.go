package request

import "lancer/common/page"

type RbacRoleListRequest struct {
	page.Request
	RoleName string `json:"role_name" form:"role_name"`
}

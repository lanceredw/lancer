package request

import "lancer/common/page"

type UserListRequest struct {
	page.Request
	UserName string `json:"user_name" form:"user_name"` //user name
	Status   int32  `json:"status" form:"status"`       //status -1.nothing  0.disable 1.enable
}

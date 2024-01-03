package response

import (
	"lancer/common/page"
	"lancer/model"
)

type RbacRoleListResponse struct {
	Code int              `json:"code"`
	Msg  string           `json:"msg"`
	Data RbacRoleListData `json:"data"`
}

type RbacRoleListData struct {
	page.Response
	List []*model.LancerRbacRole `json:"list"`
}

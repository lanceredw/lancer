package response

import "lancer/model"

type RbacMenuListResponse struct {
	Code int              `json:"code"`
	Msg  string           `json:"msg"`
	Data RbacMenuLineList `json:"data"`
}

type RbacMenuLineList struct {
	List []*RbacMenuLineData `json:"list"`
}

type RbacMenuLineData struct {
	model.LancerRbacMenu
	Children []*RbacMenuLineData `json:"children"`
}

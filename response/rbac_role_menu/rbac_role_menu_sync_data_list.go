package response

import "lancer/model"

type RbacRoleMenuSyncDataListResponse struct {
	Code int              `json:"code"`
	Msg  string           `json:"msg"`
	Data RbacMenuLineList `json:"data"`
}

type RbacMenuLineList struct {
	List map[int64][]*model.LancerRbacMenu `json:"list"`
}

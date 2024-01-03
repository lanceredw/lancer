package response

import (
	"lancer/common/page"
	"lancer/model"
)

type UserListResponse struct {
	Code int          `json:"code"`
	Msg  string       `json:"msg"`
	Data UserListData `json:"data"`
}

type UserListData struct {
	page.Response
	List []*model.LancerUser `json:"list"`
}

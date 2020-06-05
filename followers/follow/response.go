package follow

import "github.com/orderforme/user/followers/model"

type FollowList struct {
	Data         []model.Follow `json:"data"` //[]*model.Employee
	TotalRecords int            `json:"totalRecords"`
}

// DefaultFollowResponse body
type DefaultFollowResponse struct {
	Error  bool         `json:"error"`
	Follow model.Follow `json:"follow"`
}

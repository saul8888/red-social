package public

import "github.com/orderforme/user/public/model"

type PublicList struct {
	Data         []model.Public `json:"data"` //[]*model.Employee
	TotalRecords int            `json:"totalRecords"`
}

// DefaultPublicResponse body
type DefaultPublicResponse struct {
	Error  bool         `json:"error"`
	Public model.Public `json:"public"`
}

package models

// MembershipFunc ...
type MembershipFunc struct {
	BaseModel
	Name      string `json:"name"`
	Code      string `json:"code" gorm:"unique"`
	Type      string `json:"type"`
	ParamsCnt uint   `json:"params_count"`
}

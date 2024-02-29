package model

type Role struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	BaseModel
}

package model

type Role struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	BaseModel
}

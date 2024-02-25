package model

type User struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Status   int    `json:"status,omitempty"`
	Role     []Role `json:"role,omitempty"`
	BaseModel
}

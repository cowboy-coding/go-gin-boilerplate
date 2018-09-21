package models

type User struct {
	ID        string `json:"id,omitempty"`
	UserName  string `json:"user_name,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}

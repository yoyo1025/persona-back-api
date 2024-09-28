package model

type Comment struct{
	ID						int64 `json:"id,omitempty"`
	UserID				int64 `json:"user_id,omitempty"`
	PersonaID     int64 `json:"persona,omitempty"`
	Comment       string `json:"comment"`
	IsUserComment bool   `json:"is_user_comment,omitempty"`
	Good          bool	`json:"good,omitempty"`
}
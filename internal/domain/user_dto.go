package domain

type UserItemTopUser struct {
	User
	Count int64 `json:"ping_count"`
}

package models

type Permission string

const (
	Read    Permission = "r"
	Write   Permission = "w"
	Execute Permission = "x"
)

type User struct {
	ID         int        `json:"id"`
	Username   *string    `json:"username"`
	Password   *string    `json:"password"`
	IsAdmin    bool       `json:"isAdmin"`
	Permission Permission `json:"permissions"`
}

func (p Permission) IsValid() bool {
	switch p {
	case Read, Write, Execute:
		return true
	default:
		return false
	}
}

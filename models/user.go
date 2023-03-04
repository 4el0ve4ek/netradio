package models

type User struct {
	UID       int        `json:"uuid"`
	Nickname  string     `json:"name"`
	PhotoLink string     `json:"photo"`
	Lang      Lang       `json:"lang"`
	Status    UserStatus `json:"status"`
	Email     string     `json:"email"`
	Password  string     `json:"-"`
}

type UserStatus int

const (
	UserGuest UserStatus = iota
	UserRegistered
	UserAdministrator
)

type Lang string

const (
	Ru Lang = "ru"
	En Lang = "en"
)

package models

type UserSettings struct {
	From string
	To   string
}

type User struct {
	ID       string
	UserName string
	ChatID   string
	UserSettings
}

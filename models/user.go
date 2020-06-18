package models

// User model for gorm
type User struct {
	ID       string
	Name     string
	Level    string
	Room     string
	Number   string
	Plan     string
	Advisor1 string
	Advisor2 string
	RegLink  string
}

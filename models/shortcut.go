package models

type Shortcut struct {
	Id          int
	Title       string
	Description string
	Scope       string
	Shortcut    string
}

func NewShortcut(title, description, scope, shortcut string) Shortcut {
	return Shortcut{
		Title:       title,
		Description: description,
		Scope:       scope,
		Shortcut:    shortcut,
	}
}

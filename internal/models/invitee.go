package models

type Invitee struct {
	ID   string `db:"id"   json:"id"`
	Name string `db:"name" json:"name"`
}

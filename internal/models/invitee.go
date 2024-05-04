package models

type Invitee struct {
	ID        string `db:"id"`
	Name      string `db:"name"`
	Phone     int    `db:"phone"`
	Status    string `db:"status"`
	InvitedBy string `db:"invited_by"`
}

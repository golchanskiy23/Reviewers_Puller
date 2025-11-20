package entity

type User struct {
	ID       string `db:"user_id"` // TEXT в БД
	Username string `db:"username"`
	TeamName string `db:"team_name"`
	Active   bool   `db:"is_active"`
}

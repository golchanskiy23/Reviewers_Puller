package entity

import "time"

type Team struct {
	TeamName  string    `db:"team_name"`
	CreatedAt time.Time `db:"created_at"`
}

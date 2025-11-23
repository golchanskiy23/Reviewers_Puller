package postgres

import (
	"Service-for-assigning-reviewers-for-Pull-Requests/pkg/database/postgres"
	"context"
)

type Team struct {
	ID   int
	Name string
}

type TeamRepository interface {
	Add(ctx context.Context, team *Team) error
	Get(ctx context.Context, id int) (*Team, error)
}

type teamPGRepository struct {
	db *postgres.DatabaseSource
}

func NewTeamPGRepository(db *postgres.DatabaseSource) TeamRepository {
	return &teamPGRepository{db: db}
}

func (repo *teamPGRepository) Add(ctx context.Context, team *Team) error {
	return nil
}

func (repo *teamPGRepository) Get(ctx context.Context, id int) (*Team, error) {
	return nil, nil
}

/*func (r *teamPGRepository) Add(name string, users []entity.User) (*entity.Team, error) {
	/*tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`INSERT INTO teams (team_name) VALUES ($1)`, name)
	if err != nil {
		return nil, err
	}

	for _, u := range users {
		_, err = tx.Exec(`
            INSERT INTO users (user_id, username, team_name, is_active)
            VALUES ($1, $2, $3, TRUE)`,
			u.ID, u.Username, name,
		)
		if err != nil {
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &entity.Team{
		Name: name,
	}, nil
	return nil, nil
}

func (r *teamPGRepository) Get(name string) (*entity.Team, error) {
	/*t := entity.Team{}
	err := r.db.QueryRow(`
        SELECT team_name, created_at
        FROM teams
        WHERE team_name = $1`,
		name,
	).Scan(&t.Name, &t.CreatedAt)

	if err != nil {
		return nil, err
	}
	return nil, nil
}*/

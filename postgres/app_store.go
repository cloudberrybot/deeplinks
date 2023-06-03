package postgres

import (
	"fmt"

	"github.com/cloudberrybot/deeplinks"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type AppStore struct {
	*sqlx.DB
}

func (s *AppStore) App(id uuid.UUID) (deeplinks.App, error) {
	var d deeplinks.App
	if err := s.Get(&d, "SELECT * FROM apps WHERE id = $1", id); err != nil {
		return deeplinks.App{}, fmt.Errorf("error getting app: %w", err)
	}

	return d, nil
}

func (s *AppStore) Apps() ([]deeplinks.App, error) {
	var d []deeplinks.App

	if err := s.Select(&d, "SELECT * FROM apps"); err != nil {
		return []deeplinks.App{}, fmt.Errorf("error getting apps: %w", err)
	}

	return d, nil
}

func (s *AppStore) CreateApp(d *deeplinks.App) error {
	if err := s.Get(d, "INSERT INTO apps (name, description) VALUES ($1, $2) RETURNING *",
		d.Name); err != nil {
		return fmt.Errorf("error creating app: %w", err)
	}

	return nil
}

func (s *AppStore) UpdateApp(d *deeplinks.App) error {

	if err := s.Get(d, "UPDATE apps SET name = $1 WHERE id = $2 RETURNING *",
		d.Name, d.ID); err != nil {
		return fmt.Errorf("error updating app: %w", err)
	}

	return nil
}

func (s *AppStore) DeleteApp(id uuid.UUID) error {
	if _, err := s.Exec("DELETE FROM apps WHERE id = $1", id); err != nil {
		return fmt.Errorf("error deleting app: %w", err)
	}

	return nil
}

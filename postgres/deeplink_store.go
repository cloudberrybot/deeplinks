package postgres

import (
	"fmt"

	"github.com/cloudberrybot/deeplinks"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type DeepLinkStore struct {
	*sqlx.DB
}

func (s *DeepLinkStore) DeepLink(id uuid.UUID) (deeplinks.DeepLink, error) {
	var d deeplinks.DeepLink
	if err := s.Get(&d, "SELECT * FROM deeplinks WHERE id = $1", id); err != nil {
		return deeplinks.DeepLink{}, fmt.Errorf("error getting deeplink: %w", err)
	}

	return d, nil
}

func (s *DeepLinkStore) DeepLinks() ([]deeplinks.DeepLink, error) {
	var d []deeplinks.DeepLink

	if err := s.Select(&d, "SELECT * FROM deeplinks"); err != nil {
		return []deeplinks.DeepLink{}, fmt.Errorf("error getting deeplinks: %w", err)
	}

	return d, nil
}

func (s *DeepLinkStore) DeepLinksByApp(appID uuid.UUID) ([]deeplinks.DeepLink, error) {
	var d []deeplinks.DeepLink

	if err := s.Select(&d, "SELECT * FROM deeplinks WHERE app_id = $1", appID); err != nil {
		return []deeplinks.DeepLink{}, fmt.Errorf("error getting deeplinks: %w", err)
	}

	return d, nil
}

func (s *DeepLinkStore) CreateDeepLink(d *deeplinks.DeepLink) error {
	if err := s.Get(d, "INSERT INTO deeplinks (app_id, description, sharecode) VALUES ($1, $2, $3) RETURNING *",
		d.AppID, d.Description, d.UniversalLink); err != nil {
		return fmt.Errorf("error creating deeplink: %w", err)
	}

	return nil
}

func (s *DeepLinkStore) UpdateDeepLink(d *deeplinks.DeepLink) error {

	if err := s.Get(d, "UPDATE deeplinks SET app_id = $1, description = $2, sharecode = $3 WHERE id = $4 RETURNING *",
		d.AppID, d.Description, d.UniversalLink, d.ID); err != nil {
		return fmt.Errorf("error updating deeplink: %w", err)
	}

	return nil
}

func (s *DeepLinkStore) DeleteDeepLink(id uuid.UUID) error {
	if _, err := s.Exec("DELETE FROM deeplinks WHERE id = $1", id); err != nil {
		return fmt.Errorf("error deleting deeplink: %w", err)
	}

	return nil
}

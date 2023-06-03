package deeplinks

import "github.com/google/uuid"

type App struct {
	ID          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
}

type DeepLink struct {
	ID            uuid.UUID `db:"id"`
	AppID         uuid.UUID `db:"app_id"`
	Description   string    `db:"description"`
	UniversalLink uuid.UUID `db:"universal_link"`
}

type AppStore interface {
	App(id uuid.UUID) (App, error)
	Apps() ([]App, error)
	CreateApp(d *App) error
	UpdateApp(d *App) error
	DeleteApp(id uuid.UUID) error
}

type DeepLinkStore interface {
	DeepLink(id uuid.UUID) (DeepLink, error)
	DeepLinks() ([]DeepLink, error)
	DeepLinksByApp(userID uuid.UUID) ([]DeepLink, error)
	CreateDeepLink(d *DeepLink) error
	UpdateDeepLink(d *DeepLink) error
	DeleteDeepLink(id uuid.UUID) error
}

type Store interface {
	AppStore
	DeepLinkStore
}

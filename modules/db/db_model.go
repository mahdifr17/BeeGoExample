package db

import "time"

// Model is base db model
type Model struct {
	CreatedAt time.Time `json:"created_at" orm:"null;auto_now_add;type(datetime)"`
	UpdatedAt time.Time `json:"updated_at" orm:"null;auto_now;type(datetime)"`
	IsDeleted bool      `json:"is_deleted" orm:"type(bool);default(false)"`
}

// ModelActionInterface holds Model struct action
type ModelActionInterface interface {
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}

/* Implicit declaration Model implements ModelActionInterface */

// GetCreatedAt return model created at
func (m *Model) GetCreatedAt() time.Time {
	return m.CreatedAt
}

// GetUpdatedAt return model updated at
func (m *Model) GetUpdatedAt() time.Time {
	return m.UpdatedAt
}

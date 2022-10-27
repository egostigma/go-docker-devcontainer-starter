package repo

import (
	"context"
	"fmt"

	"go-docker-devcontainer-starter/internal/entity"

	"gorm.io/gorm"
)

const _defaultEntityCap = 64

type History struct {
	ID          uint   `gorm:"primaryKey" json:"id" example:"1"`
	Source      string `json:"source" example:"auto"`
	Destination string `json:"destination" example:"en"`
	Original    string `json:"original" example:"текст для перевода"`
	Translation string `json:"translation" example:"text for translation"`
}

// TranslationRepo -.
type TranslationRepo struct {
	*gorm.DB
}

// New -.
func New(sql *gorm.DB) *TranslationRepo {
	return &TranslationRepo{sql}
}

// GetHistory -.
func (r *TranslationRepo) GetHistory(ctx context.Context) ([]entity.Translation, error) {
	rows, err := r.Table("history").Rows()
	if err != nil {
		return nil, fmt.Errorf("TranslationRepo - GetHistory - r.Table('history').Rows() %w", err)
	}
	defer rows.Close()

	entities := make([]entity.Translation, 0, _defaultEntityCap)

	for rows.Next() {
		e := entity.Translation{}

		err = rows.Scan(&e.ID, &e.Source, &e.Destination, &e.Original, &e.Translation)
		if err != nil {
			return nil, fmt.Errorf("TranslationRepo - GetHistory - rows.Scan: %w", err)
		}

		entities = append(entities, e)
	}

	return entities, nil
}

// Store -.
func (r *TranslationRepo) Store(ctx context.Context, t entity.Translation) error {
	result := r.Table("history").Create(&t)

	return result.Error
}

package usecase

import (
	"context"
	"fmt"

	"go-docker-devcontainer-starter/internal/entity"
)

// TranslationUseCase -.
type TranslationUseCase struct {
	repo   TranslationRepo
	webAPI TranslationWebAPI
}

// New -.
func New(r TranslationRepo, w TranslationWebAPI) *TranslationUseCase {
	return &TranslationUseCase{
		repo:   r,
		webAPI: w,
	}
}

// History - getting translate history from store.
func (uc *TranslationUseCase) History(ctx context.Context) ([]entity.Translation, error) {
	translations, err := uc.repo.GetHistory(ctx)
	if err != nil {
		return nil, fmt.Errorf("TranslationUseCase - History - s.repo.GetHistory: %w", err)
	}

	return translations, nil
}

// Translate -.
func (uc *TranslationUseCase) Translate(ctx context.Context, t entity.Translation) (entity.Response, error) {
	translation, err := uc.webAPI.Translate(t)
	if err != nil {
		return entity.Response{
			Status:  "Translate error",
			Message: "Failed to connect to google translate",
		}, fmt.Errorf("TranslationUseCase - Translate - s.webAPI.Translate: %w", err)
	}

	err = uc.repo.Store(context.Background(), translation)
	if err != nil {
		return entity.Response{
			Status:  "Database error",
			Message: "Data save failed",
		}, fmt.Errorf("TranslationUseCase - Translate - s.repo.Store: %w", err)
	}

	return entity.Response{
		Status:  "Success",
		Message: "Data saved",
	}, nil
}

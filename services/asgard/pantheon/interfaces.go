package pantheon

import (
	"context"
	"github.com/artchitector/artchitect2/model"
)

// artPile - Куча написанных картин (репозиторий для таблицы art)
type artPile interface {
	GetNextArtID(ctx context.Context) (uint, error)
}

type ai interface {
	GenerateImage(ctx context.Context, seed uint, prompt string) ([]byte, error)
}

type bifröst interface {
	SendDrakkar(ctx context.Context, cargo model.Cargo) error
}

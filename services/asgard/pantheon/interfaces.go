package pantheon

import (
	"context"
	"github.com/artchitector/artchitect2/model"
	"image"
)

// artPile - Куча написанных картин (репозиторий для таблицы art)
type artPile interface {
	GetNextArtID(ctx context.Context) (uint, error)
	GetMaxArtID(ctx context.Context) (uint, error)
	SaveArt(ctx context.Context, artID uint, art model.Art, idea model.Idea) (model.Art, error)
	GetLastPaintTime(ctx context.Context) (uint, error) // в миллисекундах
}

// unityPile - репозиторий единств
type unityPile interface {
	Get(ctx context.Context, mask string) (model.Unity, error)
	Create(ctx context.Context, mask string, min, max uint) (model.Unity, error)
	Save(ctx context.Context, unity model.Unity) (model.Unity, error)
}

type ai interface {
	GenerateImage(ctx context.Context, seed uint, prompt string) ([]byte, error)
}

type bifröst interface {
	SendDrakkar(ctx context.Context, channel string, item interface{}) error
}

type warehouse interface {
	SaveImage(ctx context.Context, artID uint, img image.Image) error
}

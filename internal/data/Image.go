package data

import (
	"context"
	"time"

	"github.com/hossein1376/kala/internal/ent"
	entImage "github.com/hossein1376/kala/internal/ent/image"
	"github.com/hossein1376/kala/internal/structure"
)

type ImageModel struct {
	client *ent.Client
}

func (i *ImageModel) Get(id ...int) ([]*ent.Image, error) {
	return i.client.
		Image.
		Query().
		Where(entImage.IDIn(id...)).
		All(context.Background())
}

func (i *ImageModel) Insert(image *structure.Image) error {
	_, err := i.client.Image.Create().
		SetName(image.Name).
		SetCaption(image.Caption).
		SetHeight(image.Height).
		SetWidth(image.Width).
		SetSizeKB(image.Size).
		SetPath(image.Path).
		AddUserIDs(image.Owner).
		SetUploadedAt(time.Now()).
		Save(context.Background())
	return err
}

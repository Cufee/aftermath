package logic

import (
	"context"
	"errors"
	"image"

	"github.com/cufee/aftermath/internal/database"
	"github.com/cufee/aftermath/internal/database/models"
	"github.com/cufee/aftermath/internal/encoding"
	"github.com/cufee/aftermath/internal/log"
)

func UserContentToImage(content models.UserContent) (image.Image, error) {
	if content.Value == "" {
		return nil, errors.New("value is empty")
	}
	var img image.NRGBA
	err := encoding.DecodeGob([]byte(content.Value), &img)
	if err != nil {
		return nil, err
	}
	return &img, nil
}

func ImageToUserContentValue(img image.Image) ([]byte, error) {
	if img == nil {
		return nil, errors.New("image is nil")
	}
	encoded, err := encoding.EncodeGob(img)
	if err != nil {
		return nil, err
	}
	return encoded, nil
}

func GetAccountBackgroundImage(ctx context.Context, db database.Client, accountID string) (image.Image, models.UserContent, error) {
	// find who owns the account verification
	connections, err := db.FindUserConnections(ctx, database.ConnectionType(models.ConnectionTypeWargaming), database.ConnectionReferenceID(accountID), database.ConnectionVerified(true))
	if err != nil {
		return nil, models.UserContent{}, err
	}
	if len(connections) < 1 {
		return nil, models.UserContent{}, errors.New("account id does not have a verified connection")
	}
	if len(connections) > 1 {
		log.Warn().Msg("found multiple verified connections for the same wargaming account")
	}

	content, err := db.GetUserContentFromRef(ctx, connections[0].UserID, models.UserContentTypePersonalBackground)
	if err != nil {
		return nil, models.UserContent{}, err
	}
	img, err := UserContentToImage(content)
	if err != nil {
		return nil, models.UserContent{}, err
	}
	return img, content, nil
}

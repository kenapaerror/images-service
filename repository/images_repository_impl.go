package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/kenapaerror/images-service/helper"
	"github.com/kenapaerror/images-service/model/entity"
)

type ImageRepositoryImpl struct{}

func NewImageRepositoryImpl() ImageRepository {
	return &ImageRepositoryImpl{}
}

func (repository *ImageRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, image entity.Images) entity.Images {
	SQL := "INSERT INTO images (id, path, created_at, updated_at) VALUES (?,?,?,?)"

	_, err := tx.ExecContext(
		ctx,
		SQL,
		image.Id,
		image.Path,
		image.CreatedAt,
		image.UpdatedAt,
	)
	helper.PanicIfError(err)

	return image
}

func (repository *ImageRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, image entity.Images) {
	SQL := "DELETE FROM images WHERE id = ?"

	_, err := tx.ExecContext(
		ctx,
		SQL,
		image.Id,
	)
	helper.PanicIfError(err)
}

func (repository *ImageRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, imageId string) (entity.Images, error) {
	SQL := "SELECT id, path, created_at, updated_at FROM images WHERE id=?"

	rows, err := tx.QueryContext(
		ctx,
		SQL,
		imageId,
	)
	helper.PanicIfError(err)
	defer rows.Close()

	image := entity.Images{}
	if rows.Next() {
		err := rows.Scan(
			&image.Id,
			&image.Path,
			&image.CreatedAt,
			&image.UpdatedAt,
		)
		helper.PanicIfError(err)

		return image, nil
	}

	return image, errors.New("image not found")
}

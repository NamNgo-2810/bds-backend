package upldmdl

import (
	"errors"

	"bds/common"
)

const EntityName = "Upload"

type Upload struct {
	common.SQLModel `json:",inline"`
	common.Image    `json:",inline"`
}

func (Upload) TableName() string {
	return "uploads"
}

var ErrFileTooLarge = common.NewCustomError(errors.New("file is too large"), "File is too large", "ErrFileTooLarge")

func ErrFileIsNotImage(err error) *common.AppError {
	return common.NewCustomError(err, "File is not image", "ErrFileIsNotImage")
}

func ErrCannotSaveFile(err error) *common.AppError {
	return common.NewCustomError(err, "Cannot save file", "ErrCannotSaveFile")
}

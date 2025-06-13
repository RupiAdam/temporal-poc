package usecase

import (
	"fmt"
	"github.com/gin-gonic/gin"
	uuid2 "github.com/google/uuid"
	"mime/multipart"
)

type UserUsecase struct {
	ImageProcessingUsecase *ImageProcessingUsecase
}

func NewUserUsecase(imageProcessingUsecase *ImageProcessingUsecase) *UserUsecase {
	return &UserUsecase{
		ImageProcessingUsecase: imageProcessingUsecase,
	}
}

func (c *UserUsecase) UpdateProfilePicture(ctx *gin.Context, file *multipart.FileHeader) error {
	// Upload the image locally
	filename := "assets/uploads/" + file.Filename
	err := ctx.SaveUploadedFile(file, filename)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Process the image
	uuid, err := uuid2.NewUUID()
	if err != nil {
		fmt.Println(err)
		return err
	}

	outputFilename := "assets/images/" + uuid.String() + ".png"
	err = c.ImageProcessingUsecase.Resize(filename, outputFilename, 100)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

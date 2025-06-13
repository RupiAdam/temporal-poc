package usecase

import (
	"fmt"
	"github.com/h2non/bimg"
	"os"
)

type ImageProcessingUsecase struct{}

func NewImageProcessingUsecase() *ImageProcessingUsecase {
	return &ImageProcessingUsecase{}
}

func (c *ImageProcessingUsecase) Resize(filePath string, outputPath string, size int) error {
	buffer, err := bimg.Read(filePath)
	if err != nil {
		fmt.Println(os.Stderr, err)
		return err
	}

	image := bimg.NewImage(buffer)
	imageSize, err := image.Size()
	if err != nil {
		fmt.Println(os.Stderr, err)
		return err
	}

	var croppedImage []byte
	fmt.Sprintf("%dx%d", imageSize.Width, imageSize.Height)
	if imageSize.Width > imageSize.Height {
		croppedImage, err = image.CropByHeight(imageSize.Height)
		if err != nil {
			fmt.Println(os.Stderr, err)
			return err
		}
	} else {
		croppedImage, err = image.CropByWidth(imageSize.Width)
		if err != nil {
			fmt.Println(os.Stderr, err)
			return err
		}
	}

	newImage, err := bimg.NewImage(croppedImage).Resize(size, size)
	if err != nil {
		fmt.Println(os.Stderr, err)
		return err
	}

	err = bimg.Write(outputPath, newImage)
	if err != nil {
		fmt.Println(os.Stderr, err)
		return err
	}
	return nil
}

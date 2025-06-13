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

func (c *ImageProcessingUsecase) Resize(filePath string, outputPath string, targetSize int) error {
	buffer, err := bimg.Read(filePath)
	if err != nil {
		fmt.Println(os.Stderr, err)
		return err
	}

	image := bimg.NewImage(buffer)
	meta, err := image.Metadata()
	if err != nil {
		return fmt.Errorf("unable to read metadata: %w", err)
	}
	w, h := meta.Size.Width, meta.Size.Height

	// 2. Determine square side and offset to center‐crop
	var side, left, top int
	if w < h {
		side = w
		left = 0
		top = (h - w) / 2
	} else {
		side = h
		left = (w - h) / 2
		top = 0
	}

	// 3. Crop to square
	cropOpts := bimg.Options{
		Crop:    true,
		Width:   side,
		Height:  side,
		Top:     top,
		Left:    left,
		Quality: 95,
	}
	squareBuf, err := image.Process(cropOpts)
	if err != nil {
		return fmt.Errorf("crop failed: %w", err)
	}

	newImage, err := bimg.NewImage(squareBuf).Process(bimg.Options{
		Width:   targetSize,
		Height:  targetSize,
		Quality: 95,
	})
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

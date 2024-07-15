package infrastructure

import (
	"fmt"
	"os"

	"github.com/google/uuid"
)

type ImageStore struct {
	imageFolderPath string
}

func NewImageStore(imageFolderPath string) *ImageStore {
	imgStore := &ImageStore{
		imageFolderPath: imageFolderPath,
	}
	return imgStore
}

func (is *ImageStore) SaveImage(imageBytes []byte) (string, error) {
	photoUuid, err := uuid.NewV7()
	if err != nil {
		return "", err
	}
	imagePath := fmt.Sprintf("%s/%s.jpg", is.imageFolderPath, photoUuid)
	file, err := os.Create(imagePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	file.Write(imageBytes)

	return imagePath, nil
}

func (is *ImageStore) SaveImageByPath(pathToSave string, imageBytes []byte) (string, error) {
	photoUuid, err := uuid.NewV7()
	if err != nil {
		return "", err
	}
	imagePath := fmt.Sprintf("%s/%s.jpg", pathToSave, photoUuid)
	file, err := os.Create(imagePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	file.Write(imageBytes)

	return imagePath, nil
}

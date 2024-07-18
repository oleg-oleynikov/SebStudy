package infrastructure

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/uuid"
)

type ImageStore struct {
	imageFolderPath string
}

func NewImageStore(imageFolderPath string) *ImageStore {
	_, err := os.Stat(imageFolderPath)
	if os.IsNotExist(err) {
		os.MkdirAll(imageFolderPath, 0755)
		if err != nil {
			log.Fatalf("failed to create folder for imageStore")
			return nil
		}
	}

	imgStore := &ImageStore{
		imageFolderPath: imageFolderPath,
	}
	return imgStore
}

func (is *ImageStore) SaveImage(imageBytes []byte) (string, error) {
	imageUuid, err := uuid.NewV7()
	if err != nil {
		return "", err
	}

	imagePath := fmt.Sprintf("%s/%s.jpg", is.imageFolderPath, imageUuid)
	file, err := os.Create(imagePath)

	if err != nil {
		return "", err
	}
	defer file.Close()
	file.Write(imageBytes)

	return imagePath, nil
}

func (is *ImageStore) GetImageBytes(imagePath string) ([]byte, error) {
	if !strings.HasSuffix(imagePath, ".jpg") {
		return nil, fmt.Errorf("image must has a suffix .jpg")
	}

	imageBytes, err := os.ReadFile(imagePath)
	if err != nil {
		return nil, err
	}

	return imageBytes, nil
}

func (is *ImageStore) DeleteImageByPath(imagePath string) error {
	if !strings.HasSuffix(imagePath, ".jpg") {
		return fmt.Errorf("image must has a suffix .jpg")
	}

	if err := os.Remove(imagePath); err != nil {
		return err
	}

	return nil
}

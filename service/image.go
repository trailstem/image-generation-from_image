package service

import (
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"os"

	"github.com/nfnt/resize"
)

func ResizeImage(inputPath string, width, height uint) (string, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return "", fmt.Errorf("failed to decode image: %w", err)
	}

	resizedImg := resize.Resize(width, height, img, resize.Lanczos3)

	outputPath := "resized_profile.png"
	outFile, err := os.Create(outputPath)
	if err != nil {
		return "", fmt.Errorf("failed to create resized file: %w", err)
	}
	defer outFile.Close()

	jpeg.Encode(outFile, resizedImg, nil)
	return outputPath, nil
}

func EncodeImageToBase64(filePath string) (string, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}
	return base64.StdEncoding.EncodeToString(file), nil
}

// SaveGeneratedImage は生成された画像をファイルとして保存します
func SaveGeneratedImage(data []byte, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("failed to write data to file: %w", err)
	}

	return nil
}

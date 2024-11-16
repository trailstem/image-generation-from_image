package service

import (
	"encoding/base64"
	"fmt"
	"image/color"
	"image/png"
	"io"
	"io/ioutil"
	"os"
)

// 画像ファイルを Base64 エンコードし、MIME タイプのプレフィックスを追加
func EncodeImageToBase64WithPrefix(imagePath string, format string) (string, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	imgBytes, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	encoded := base64.StdEncoding.EncodeToString(imgBytes)

	// MIME タイプのプレフィックスを追加
	var prefix string
	if format == "jpeg" {
		prefix = "data:image/jpeg;base64,"
	} else if format == "png" {
		prefix = "data:image/png;base64,"
	} else {
		return "", fmt.Errorf("unsupported format: %s", format)
	}

	return prefix + encoded, nil
}

func SaveBase64Image(response string, filePath string) error {
	data, err := base64.StdEncoding.DecodeString(response)
	if err == nil {
		return ioutil.WriteFile(filePath, data, 0644)
	}
	return ioutil.WriteFile(filePath, []byte(response), 0644)
}

// マスク画像を整数配列に変換する関数
func ConvertMaskToIntArray(filePath string) ([]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	maskArray := make([]int, width*height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			grayColor := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			maskArray[y*width+x] = int(grayColor.Y)
		}
	}
	return maskArray, nil
}

func FormatBase64Image(encodedImage, format string) string {
	if format == "jpeg" {
		return fmt.Sprintf("data:image/jpeg;base64,%s", encodedImage)
	} else if format == "png" {
		return fmt.Sprintf("data:image/png;base64,%s", encodedImage)
	}
	return encodedImage
}

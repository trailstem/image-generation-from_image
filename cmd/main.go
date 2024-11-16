package main

import (
	"encoding/base64"
	"fmt"

	"github.com/trailstem/image-generation-from-image/config"
	"github.com/trailstem/image-generation-from-image/service"
	"github.com/trailstem/image-generation-from-image/usecase"
)

func main() {
	// 環境変数の読み込み
	env, err := config.LoadEnv()
	if err != nil {
		fmt.Println("Error loading environment variables:", err)
		return
	}

	// 入力画像のリサイズとエンコード
	imagePath := "./profile.png"
	resizedImagePath, err := service.ResizeImage(imagePath, 512, 512)
	if err != nil {
		fmt.Println("Error resizing image:", err)
		return
	}

	encodedImage, err := service.EncodeImageToBase64(resizedImagePath)
	if err != nil {
		fmt.Println("Error encoding image:", err)
		return
	}

	// マスク画像のリサイズとエンコード
	maskPath := "./mask.png"
	resizedMaskPath, err := service.ResizeImage(maskPath, 512, 512)
	if err != nil {
		fmt.Println("Error resizing mask:", err)
		return
	}

	encodedMask, err := service.EncodeImageToBase64(resizedMaskPath)
	if err != nil {
		fmt.Println("Error encoding mask:", err)
		return
	}

	// 画像生成
	resultImage, err := usecase.GenerateImageWithInpainting(env, encodedImage, encodedMask)
	if err != nil {
		fmt.Println("Error generating image:", err)
		return
	}

	// 生成された画像データ (Base64エンコード) をデコードしてバイナリデータに変換
	imageData, err := base64.StdEncoding.DecodeString(resultImage)
	if err != nil {
		fmt.Println("Error decoding Base64 image:", err)
		return
	}

	// 生成された画像を保存
	err = service.SaveGeneratedImage(imageData, "./output.png")
	if err != nil {
		fmt.Println("Error saving generated image:", err)
		return
	}
	fmt.Println("Image generated successfully: output.png")
}

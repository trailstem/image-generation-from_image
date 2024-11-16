package usecase

import (
	"github.com/trailstem/image-generation-from-image/config"
	"github.com/trailstem/image-generation-from-image/model"
	"github.com/trailstem/image-generation-from-image/repository"
	"github.com/trailstem/image-generation-from-image/service"
)

// func GenerateImage(env *config.Env, imageB64 string) (string, error) {
// 	requestBody := model.Img2ImgRequest{
// 		Prompt:         "Add a black baseball cap on the person's head, keeping the original style of the image.",
// 		NegativePrompt: "cartoon, unrealistic, distorted, low resolution, blurry",
// 		Height:         512,
// 		Width:          512,
// 		ImageB64:       imageB64,
// 		NumSteps:       20,
// 		Strength:       0.6, // 元の画像をもう少し保持
// 		Guidance:       9.0, // プロンプトへの忠実度を上げる
// 	}

// 	return repository.GenerateImageFromAPI(env.AccountID, env.ApiToken, requestBody)
// }


func GenerateImageWithInpainting(env *config.Env, imageB64 string, maskPath string) (string, error) {
	// マスク画像を整数配列に変換
	mask, err := service.ConvertMaskToIntArray(maskPath)
	if err != nil {
		return "", err
	}

	requestBody := model.InpaintingRequest{
		Prompt:         "Add a black baseball cap on the person's head.",
		NegativePrompt: "cartoon, unrealistic, distorted, low resolution, blurry",
		Height:         512,
		Width:          512,
		ImageB64:       imageB64,
		Mask:           mask,
		NumSteps:       20,
		Strength:       0.5,
		Guidance:       8.0,
	}

	return repository.GenerateImageFromAPI(env.AccountID, env.ApiToken, requestBody)
}
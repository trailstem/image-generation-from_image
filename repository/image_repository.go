package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/trailstem/image-generation-from-image/model"
)

// 画像生成APIを呼び出す関数
// func GenerateImageFromAPI(accountID, apiToken string, requestBody model.Img2ImgRequest) (string, error) {
// 	url := fmt.Sprintf("https://api.cloudflare.com/client/v4/accounts/%s/ai/run/@cf/runwayml/stable-diffusion-v1-5-img2img", accountID)

// 	// リクエストボディを JSON にエンコード
// 	jsonBody, err := json.Marshal(requestBody)
// 	if err != nil {
// 		return "", err
// 	}

// 	// HTTP リクエストの設定
// 	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
// 	if err != nil {
// 		return "", err
// 	}
// 	req.Header.Set("Authorization", "Bearer "+apiToken)
// 	req.Header.Set("Content-Type", "application/json")

// 	// HTTP クライアントでリクエストを送信
// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer resp.Body.Close()

// 	// レスポンスの Content-Type を確認
// 	contentType := resp.Header.Get("Content-Type")

// 	// JSON レスポンスの場合
// 	if contentType == "application/json" {
// 		var response model.Img2ImgResponse
// 		body, err := ioutil.ReadAll(resp.Body)
// 		if err != nil {
// 			return "", err
// 		}
// 		if err := json.Unmarshal(body, &response); err != nil {
// 			return "", err
// 		}
// 		return response.ImageB64, nil
// 	}

// 	// 画像 (image/png) の場合
// 	if contentType == "image/png" {
// 		imagePath := "output.png"
// 		err := savePNGImage(resp, imagePath)
// 		if err != nil {
// 			return "", err
// 		}
// 		fmt.Printf("Image saved to: %s\n", imagePath)
// 		return "Image saved successfully", nil
// 	}

// 	// その他のレスポンス形式
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return "", err
// 	}
// 	return "", fmt.Errorf("unexpected content type: %s, body: %s", contentType, string(body))
// }

// // PNG 画像を保存する関数
// func savePNGImage(resp *http.Response, filePath string) error {
// 	imageData, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return err
// 	}
// 	return ioutil.WriteFile(filePath, imageData, 0644)
// }

func GenerateImageFromAPI(accountID, apiToken string, requestBody model.InpaintingRequest) (string, error) {
	url := fmt.Sprintf("https://api.cloudflare.com/client/v4/accounts/%s/ai/run/@cf/runwayml/stable-diffusion-v1-5-inpainting", accountID)

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+apiToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

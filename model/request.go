package model

type Img2ImgRequest struct {
	Prompt         string  `json:"prompt"`
	NegativePrompt string  `json:"negative_prompt,omitempty"`
	Height         int     `json:"height"`
	Width          int     `json:"width"`
	ImageB64       string  `json:"image_b64"`
	NumSteps       int     `json:"num_steps"`
	Strength       float64 `json:"strength"`
	Guidance       float64 `json:"guidance"`
	Seed           int     `json:"seed,omitempty"`
}


type InpaintingRequest struct {
	Prompt         string  `json:"prompt"`
	NegativePrompt string  `json:"negative_prompt,omitempty"`
	Height         int     `json:"height"`
	Width          int     `json:"width"`
	ImageB64       string  `json:"image_b64"`
	Mask           []int   `json:"mask"`
	NumSteps       int     `json:"num_steps"`
	Strength       float64 `json:"strength"`
	Guidance       float64 `json:"guidance"`
	Seed           int     `json:"seed,omitempty"`
}

type Img2ImgResponse struct {
	ImageB64 string `json:"image_b64"`
}

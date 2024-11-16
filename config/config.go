package config

import (
	"fmt"
	"os"
)

type Env struct {
	AccountID string
	ApiToken  string
}

func LoadEnv() (*Env, error) {
	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	apiToken := os.Getenv("CLOUDFLARE_API_TOKEN")

	if accountID == "" || apiToken == "" {
		return nil, fmt.Errorf("missing environment variables")
	}

	return &Env{
		AccountID: accountID,
		ApiToken:  apiToken,
	}, nil
}

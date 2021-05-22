package zoom

import (
	"context"
	"encoding/json"
	"fmt"
)

const (
	authPath = "oauth/token"
)

type (
	AuthResponse struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int    `json:"expires_in"`
	}
)

func (c *Client) auth(ctx context.Context) error {
	body := struct{ Authorization string }{c.token}
	payload, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("failed to parse request body: %w", err)
	}
	res, err := c.rClient.R().
		SetContext(ctx).
		SetHeader("Accept", "application/json").
		SetPathParams(map[string]string{"grant_type": "client_credentials"}).
		SetBody(payload).Get(authPath)
	if err != nil {
		return fmt.Errorf("failed in request to '%s': %w", authPath, err)
	}
	var authRes AuthResponse
	if err := json.Unmarshal(res.Body(), &authRes); err != nil {
		return fmt.Errorf("failed to parse response body: %w", err)
	}

	c.rClient.SetAuthToken(authRes.AccessToken)
	return nil
}

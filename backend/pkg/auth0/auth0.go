package auth0

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Auth0 interface {
	GetUserInfo(accessToken string) (*Userinfo, error)
}

type auth0 struct {
	domain string
}

type Userinfo struct {
	Email     string `json:"email"`
	Name      string `json:"name"`
	UpdatedAt string `json:"updated_at"`
}

func NewAuth0(domain string) Auth0 {
	return &auth0{
		domain: domain,
	}
}

func (a *auth0) GetUserInfo(accessToken string) (*Userinfo, error) {
	url := fmt.Sprintf("%s/userinfo", a.BaseURL())
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if !(200 <= resp.StatusCode && resp.StatusCode <= 299) {
		return nil, fmt.Errorf("HTTP Error (%d) : %s", resp.StatusCode, body)
	}

	var u Userinfo
	if err := json.Unmarshal(body, &u); err != nil {
		return nil, err
	}

	return &u, nil
}

func (a auth0) BaseURL() string {
	return fmt.Sprintf("https://%s", a.domain)
}

func (a auth0) JWKsURL() string {
	return fmt.Sprintf("%s/.well-known/jwks.json", a.BaseURL())
}

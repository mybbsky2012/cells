package hydra

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/pydio/cells/common"
	"github.com/pydio/cells/common/config"
	"github.com/pydio/cells/common/micro"
	"github.com/pydio/cells/common/proto/auth"
	"golang.org/x/oauth2"
)

var (
	hydraBaseURL  = config.Get("defaults", "urlInternal").String("")
	hydraAdminURL = hydraBaseURL + "/oidc-admin"
	hydraOAuthURL = hydraBaseURL + "/oidc/oauth2"
)

type ConsentResponse struct {
	Challenge                    string   `json:"challenge"`
	Skip                         bool     `json:"skip"`
	Subject                      string   `json:"subject"`
	RequestedScope               []string `json:"requested_scope"`
	RequestedAccessTokenAudience []string `json:"requested_access_token_audience"`
}

type LogoutResponse struct {
	RequestURL  string `json:"request_url"`
	RPInitiated bool   `json:"rp_initiated"`
	SID         string `json:"sid"`
	Subject     string `json:"subject"`
}

type RedirectResponse struct {
	RedirectTo string `json:"redirect_to"`
}

type TokenResponse struct {
	IDToken      string `json:"id_token"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
}

func GetLogin(challenge string) (*auth.GetLoginResponse, error) {
	c := auth.NewLoginProviderClient(common.SERVICE_GRPC_NAMESPACE_+common.SERVICE_OAUTH, defaults.NewClient())
	resp, err := c.GetLogin(context.Background(), &auth.GetLoginRequest{
		Challenge: challenge,
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func CreateLogin() (string, error) {
	c := auth.NewLoginProviderClient(common.SERVICE_GRPC_NAMESPACE_+common.SERVICE_OAUTH, defaults.NewClient())
	resp, err := c.CreateLogin(context.Background(), &auth.CreateLoginRequest{})
	if err != nil {
		return "", err
	}

	return resp.Challenge, nil
}

func AcceptLogin(challenge string, subject string) (*RedirectResponse, error) {
	c := auth.NewLoginProviderClient(common.SERVICE_GRPC_NAMESPACE_+common.SERVICE_OAUTH, defaults.NewClient())
	_, err := c.AcceptLogin(context.Background(), &auth.AcceptLoginRequest{
		Challenge: challenge,
		Subject:   subject,
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func RejectLogin(challenge string, body interface{}) (*RedirectResponse, error) {
	resp := new(RedirectResponse)

	if err := put("login", "reject", challenge, body, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func GetConsent(challenge string) (*auth.GetConsentResponse, error) {
	c := auth.NewConsentProviderClient(common.SERVICE_GRPC_NAMESPACE_+common.SERVICE_OAUTH, defaults.NewClient())
	resp, err := c.GetConsent(context.Background(), &auth.GetConsentRequest{
		Challenge: challenge,
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func CreateConsent(loginChallenge string) (string, error) {

	login, err := GetLogin(loginChallenge)
	if err != nil {
		return "", err
	}

	c := auth.NewConsentProviderClient(common.SERVICE_GRPC_NAMESPACE_+common.SERVICE_OAUTH, defaults.NewClient())
	resp, err := c.CreateConsent(context.Background(), &auth.CreateConsentRequest{
		LoginChallenge: login.Challenge,
	})
	if err != nil {
		return "", err
	}

	return resp.Challenge, nil
}

func AcceptConsent(challenge string) (*RedirectResponse, error) {
	c := auth.NewConsentProviderClient(common.SERVICE_GRPC_NAMESPACE_+common.SERVICE_OAUTH, defaults.NewClient())
	_, err := c.AcceptConsent(context.Background(), &auth.AcceptConsentRequest{
		Challenge: challenge,
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func RejectConsent(challenge string, body interface{}) (*RedirectResponse, error) {
	resp := new(RedirectResponse)

	if err := put("consent", "reject", challenge, body, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func CreateAuthCode(consentChallenge string) (string, error) {
	consent, err := GetConsent(consentChallenge)
	if err != nil {
		return "", err
	}

	c := auth.NewAuthCodeProviderClient(common.SERVICE_GRPC_NAMESPACE_+common.SERVICE_OAUTH, defaults.NewClient())
	resp, err := c.CreateAuthCode(context.Background(), &auth.CreateAuthCodeRequest{
		ConsentChallenge:  consent.Challenge,
		LoginSessionID:    consent.LoginSessionID,
		Subject:           consent.Subject,
		SubjectIdentifier: consent.SubjectIdentifier,
		ClientID:          consent.ClientID,
	})
	if err != nil {
		return "", err
	}

	return resp.Code, nil
}

func GetLogout(challenge string) (*LogoutResponse, error) {
	resp := new(LogoutResponse)

	if err := get("logout", challenge, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func AcceptLogout(challenge string, body interface{}) (*RedirectResponse, error) {
	resp := new(RedirectResponse)

	if err := put("logout", "accept", challenge, body, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func RejectLogout(challenge string, body interface{}) (*RedirectResponse, error) {
	resp := new(RedirectResponse)

	if err := put("logout", "reject", challenge, body, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func get(flow string, challenge string, res interface{}) error {
	cli := &http.Client{}

	req, err := http.NewRequest("GET", hydraAdminURL+"/oauth2/auth/requests/"+flow, nil)

	q := req.URL.Query()
	q.Add(flow+"_challenge", challenge)
	req.URL.RawQuery = q.Encode()

	resp, err := cli.Do(req)
	if err != nil {
		return err
	}

	if err := json.NewDecoder(resp.Body).Decode(res); err != nil {
		return err
	}

	return nil
}

func put(flow string, action string, challenge string, body interface{}, res interface{}) error {
	cli := &http.Client{}

	data, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", hydraAdminURL+"/oauth2/auth/requests/"+flow+"/"+action, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")

	q := req.URL.Query()
	q.Add(flow+"_challenge", challenge)
	req.URL.RawQuery = q.Encode()

	resp, err := cli.Do(req)
	if err != nil {
		return err
	}

	if err := json.NewDecoder(resp.Body).Decode(res); err != nil {
		return err
	}

	return nil
}

func Exchange(code string) (*oauth2.Token, error) {

	c := auth.NewAuthCodeExchangerClient(common.SERVICE_GRPC_NAMESPACE_+common.SERVICE_OAUTH, defaults.NewClient())
	resp, err := c.Exchange(context.Background(), &auth.ExchangeRequest{
		Code: code,
	})
	if err != nil {
		return nil, err
	}

	token := &oauth2.Token{
		AccessToken:  resp.AccessToken,
		RefreshToken: resp.RefreshToken,
		Expiry:       time.Now().Add(time.Duration(resp.Expiry) * time.Second),
	}

	token = token.WithExtra(map[string]interface{}{
		"expires_in": resp.Expiry,
		"id_token":   resp.IDToken,
	})

	return token, nil
}

func Refresh(refreshToken string) (*TokenResponse, error) {
	c := auth.NewAuthTokenRefresherClient(common.SERVICE_GRPC_NAMESPACE_+common.SERVICE_OAUTH, defaults.NewClient())
	resp, err := c.Refresh(context.Background(), &auth.RefreshTokenRequest{
		RefreshToken: refreshToken,
	})
	if err != nil {
		return nil, err
	}

	return &TokenResponse{
		AccessToken:  resp.AccessToken,
		IDToken:      resp.IDToken,
		RefreshToken: resp.RefreshToken,
		ExpiresIn:    resp.Expiry,
	}, nil
}
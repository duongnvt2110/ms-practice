package auth

import (
	"auth-service/pkg/util/goauth2"
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

// Public
func (h *authHandler) OauthGoogleLogin(c echo.Context) error {
	// Create oauthState cookie
	oauthState := h.generateStateOauthCookie(c)
	// spew.Dump(c.Scheme())
	// spew.Dump(c.Echo().Reverse("oauth.callback"))
	oauthCfg := h.getOauthConfig(c.Scheme(), c.Echo().Reverse("oauth.callback"))
	u := oauthCfg.AuthCodeURL(oauthState)
	return c.Redirect(http.StatusTemporaryRedirect, u)
}

func (h *authHandler) OauthGoogleCallback(c echo.Context) error {
	oauthState, _ := c.Cookie("oauthstate")
	if c.FormValue("state") != oauthState.Value {
		log.Println("invalid oauth google state")
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	}

	oauthCfg := h.getOauthConfig(c.Request().URL.Scheme, c.Echo().Reverse("oauth.callback"))
	data, err := h.getUserDataFromGoogle(oauthCfg, c.FormValue("code"))
	if err != nil {
		log.Println(err.Error())
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	}

	// GetOrCreate User in your db.
	// Redirect or response with a token.
	// More code .....
	spew.Dump(data)
	return nil
}

// Private
func (h *authHandler) generateStateOauthCookie(c echo.Context) string {
	var expiration = time.Now().Add(365 * 24 * time.Hour)

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := &http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	c.SetCookie(cookie)

	return state
}

func (h *authHandler) getUserDataFromGoogle(oauthCfg *oauth2.Config, code string) ([]byte, error) {
	// Use code to get token and get user info from Google.

	token, err := oauthCfg.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}
	response, err := http.Get(h.cfg.Google.OauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed read response: %s", err.Error())
	}
	return contents, nil
}

func (h *authHandler) getOauthConfig(scheme string, path string) *oauth2.Config {
	reUrl := url.URL{
		Scheme: scheme,
		Host:   h.cfg.App.Host,
		Path:   path,
	}
	spew.Dump(reUrl)
	return goauth2.GetOauth2Config(h.cfg, reUrl.String())
}

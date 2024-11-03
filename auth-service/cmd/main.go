package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOauthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:3001/auth/google/callback",
	ClientID:     "{CLIENT_ID}",
	ClientSecret: "{CLIENT_SECRET}",
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func main() {
	e := echo.New()
	e.GET("/auth/google/login", oauthGoogleLogin)
	e.GET("/auth/google/callback", oauthGoogleCallback)
	e.Logger.Fatal(e.Start(":3000"))
}

func oauthGoogleLogin(c echo.Context) error {

	// Create oauthState cookie
	oauthState := generateStateOauthCookie(c)

	/*
		AuthCodeURL receive state that is a token to protect the user from CSRF attacks. You must always provide a non-empty string and
		validate that it matches the the state query parameter on your redirect callback.
	*/

	u := googleOauthConfig.AuthCodeURL(oauthState)
	return c.Redirect(http.StatusTemporaryRedirect, u)
}

func oauthGoogleCallback(c echo.Context) error {
	// Read oauthState from Cookie
	oauthState, _ := c.Cookie("oauthstate")

	if c.FormValue("state") != oauthState.Value {
		log.Println("invalid oauth google state")
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	}

	data, err := getUserDataFromGoogle(c.FormValue("code"))
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

func generateStateOauthCookie(c echo.Context) string {
	var expiration = time.Now().Add(365 * 24 * time.Hour)

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := &http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	c.SetCookie(cookie)

	return state
}

func getUserDataFromGoogle(code string) ([]byte, error) {
	// Use code to get token and get user info from Google.

	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}
	response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
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

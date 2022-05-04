package model

import (
	"golang.org/x/oauth2"
	envConfig "zhaoxin-api/config"
)

var (
	GoogleOAuth2Config = &oauth2.Config{
		ClientID:     envConfig.Config("GOOGLE_CLIENT_ID"),
		ClientSecret: envConfig.Config("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  envConfig.Config("BACKEND_API_URL") + "/v1/auth/google/redirect",
		Endpoint: 	  oauth2.Endpoint{
			AuthURL:  "https://accounts.google.com/o/oauth2/auth",
			TokenURL: "https://accounts.google.com/o/oauth2/token",
		},
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email"},
	}
	OAuthStateStr = "random"
)
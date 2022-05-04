package model

import (
	envConfig "zhaoxin-api/config"
)

var (
	ClientID 	 string
	ClientSecret string
)

func init() {
	ClientID = envConfig.Config("GITHUB_CLIENT_ID")
	ClientSecret = envConfig.Config("GITHUB_CLIENT_SECRET")
	//storage.Sqlite3Migrate(&OAuthUser{},&StaffCache{},&WorkUp{},&StaffReplicate{})
}

type OAuthAccessResponse struct {
	AccessToken 	 string `json:"access_token"`
	TokenType 		 string `json:"token_type"`
	Scope 			 string `json:"scope"`
	Error 			 string `json:"error"`
	ErrorDescription string `json:"error_description"`
	ErrorURI 		 string `json:"error_uri"`
}
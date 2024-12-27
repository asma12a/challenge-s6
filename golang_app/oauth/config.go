package oauth

import (
	"github.com/asma12a/challenge-s6/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var GoogleOAuthConfig = &oauth2.Config{
	ClientID:     config.Env.ClientID,
	ClientSecret: config.Env.ClientSecret,
	RedirectURL:  "http://localhost:" + config.Env.APIPort + "/auth/google/callback",
	Scopes: []string{
		"https://www.googleapis.com/auth/userinfo.profile",
		"https://www.googleapis.com/auth/userinfo.email",
	},
	Endpoint: google.Endpoint,
}

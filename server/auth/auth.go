package auth

import (
	"context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var GoogleOAuthConfig *oauth2.Config

func InitGoogleOAuth(clientID, clientSecret, redirectURL string) {
	GoogleOAuthConfig = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes: []string{"https://www.googleapis.com/auth/calendar.readonly"},
		Endpoint:     google.Endpoint,
	}
}

func GetGoogleLoginURL(state string) string {
	return GoogleOAuthConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)
}

func ExchangeCodeForToken(code string) (*oauth2.Token, error) {
	return GoogleOAuthConfig.Exchange(context.Background(), code)
}
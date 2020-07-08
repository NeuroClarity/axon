package auth

import (
	"context"
	"log"

	oidc "github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

type Authenticator struct {
	Provider *oidc.Provider
	Config   oauth2.Config
	Ctx      context.Context
}

func NewAuthenticator() (*Authenticator, error) {
	ctx := context.Background()

	provider, err := oidc.NewProvider(ctx, "https://dev-q7h0r088.us.auth0.com/")
	if err != nil {
		log.Printf("Failed to get Auth0 provider: %v.\n", err)
		return nil, err
	}

	// TODO: config hardcoded values
	conf := oauth2.Config{
		ClientID:     "ih5Ms51935CplO4inqwN1RL6mJxC5LMH",
		ClientSecret: "yo7tQGafvYT5xTvMXUhBgp8l9_ph5LYGCIPR1U7GqBv1XbT8R42mvCirTOzMeAKf",
		RedirectURL:  "http://localhost:8000/api/reviewer/callback",
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile"},
	}

	return &Authenticator{
		Provider: provider,
		Config:   conf,
		Ctx:      ctx,
	}, nil
}

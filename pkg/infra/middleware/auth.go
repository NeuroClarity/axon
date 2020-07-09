package middleware

import (
	"encoding/json"
	"errors"
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
)

func NewJWTMiddleware() *jwtmiddleware.JWTMiddleware {
	return jwtmiddleware.New(jwtmiddleware.Options{

		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			// Verify 'aud' claim (Audience).
			audience := "http://axon.neuroclarity.ai"
			checkAudience := token.Claims.(jwt.MapClaims).VerifyAudience(audience, false)
			if !checkAudience {
				return token, errors.New("Invalid audience field in API header.")
			}

			// Verify 'iss' claim (Issuer).
			iss := "https://dev-q7h0r088.us.auth0.com/"
			checkISS := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
			if !checkISS {
				return token, errors.New("Invalid issuer field in API header.")
			}

			cert, err := getPemCert(token)
			if err != nil {
				return nil, err
			}

			result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
			if err != nil {
				return nil, err
			}
			return result, nil
		},

		SigningMethod: jwt.SigningMethodRS256,
	})
}

// Grabs the JSON Web Key Set and returns a certificate with the Auth0 public
// key
func getPemCert(token *jwt.Token) (string, error) {
	cert := ""

	resp, err := http.Get("https://dev-q7h0r088.us.auth0.com/.well-known/jwks.json")
	if err != nil {
		return cert, err
	}
	defer resp.Body.Close()

	// JSONWebKeys are defined by Auth0, part of their handshake.
	var jwks = Jwks{}
	err = json.NewDecoder(resp.Body).Decode(&jwks)
	if err != nil {
		return cert, err
	}

	for k := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		return cert, errors.New("Unable to find appropriate key.")
	}

	return cert, nil
}

type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

type Response struct {
	Message string `json:"message"`
}

package notify

import (
	"net/http"
	"net/url"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// Configuration of the Notifications Go Client.
type Configuration struct {
	APIKey     []byte
	BaseURL    *url.URL
	Claims     *jwt.StandardClaims
	HTTPClient *http.Client
	ServiceID  string
}

// Authenticate a JWT token. JwtTokenCreator uses HMAC-SHA256 signature, by default.
func (c *Configuration) Authenticate(secret []byte) (*string, error) {
	if c.Claims == nil {
		c.Claims = &jwt.StandardClaims{
			Issuer:   c.ServiceID,
			IssuedAt: time.Now().Unix(),
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c.Claims)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}

package token

import (
	"context"
	"time"

	"github.com/namo-io/go-kit/log"

	jwt "github.com/dgrijalva/jwt-go"
)

// Token jwt token
type Token string

func (c Token) String() string {
	return string(c)
}

// Configuration jwt token configuration
type Configuration struct {
	Secret                   string
	ExpirationTimeMilisecond int
}

// TokenStore jwt token issue
type TokenStore struct {
	Secret         []byte
	ExpirationTime time.Duration
}

// BaseClaims base claims
type BaseClaims struct {
	jwt.StandardClaims
}

// New create token store
func New(ctx context.Context, cfg *Configuration) *TokenStore {
	logger := log.New().WithContext(ctx)

	if cfg == nil {
		logger.Warn("not found jwt configuration")
		return nil
	}

	if len(cfg.Secret) <= 10 {
		logger.Warn("jwt token secret too short")
	}

	return &TokenStore{
		Secret:         []byte(cfg.Secret),
		ExpirationTime: time.Duration(cfg.ExpirationTimeMilisecond) * time.Millisecond,
	}
}

// GetBaseClaims get baseclaims
func (s *TokenStore) GetBaseClaims() *BaseClaims {
	return &BaseClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(s.ExpirationTime).Unix(),
		},
	}
}

// IssueToken issue token
func (s *TokenStore) IssueToken(ctx context.Context, claims jwt.Claims) (*Token, error) {
	logger := log.New().WithContext(ctx)

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenString, err := token.SignedString([]byte(s.Secret))
	if err != nil {
		logger.Error(err)
		return nil, nil
	}

	retn := Token(tokenString)
	return &retn, nil
}

// VertifyToken vertify token
func (s *TokenStore) VertifyToken(ctx context.Context, token string, claims jwt.Claims) error {
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.Secret), nil
	})

	return err
}

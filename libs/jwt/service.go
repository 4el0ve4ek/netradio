package jwt

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"netradio/models"

	"github.com/golang-jwt/jwt/v4"
)

const (
	authorizationHeader    = "Authorization"
	authorizationJWTPrefix = "Bearer"

	jwtExpire = time.Hour * 24 * 10
)

var (
	NoAuthorizationHeader      = errors.New("no authorization header provided")
	InvalidAuthorizationHeader = errors.New("invalid authorization header")
	CannotParseUID             = errors.New("unable to parse uid from jwt")
)

type Verificator interface {
	GetUIDFromHeader(header http.Header) (int, error)
	AddUIDToHeader(header http.Header, user models.User) error
	DeleteAuth(header http.Header)
}

func NewVerificator(config Config) *authVerificator {
	return &authVerificator{
		config: config,
	}
}

type authVerificator struct {
	config Config
}

func (a *authVerificator) GetUIDFromHeader(header http.Header) (int, error) {
	headerValue := header.Get(authorizationHeader)
	if headerValue == "" {
		return 0, NoAuthorizationHeader
	}

	headerParts := strings.Split(headerValue, " ")
	if len(headerParts) != 2 || headerParts[0] != authorizationJWTPrefix {
		return 0, InvalidAuthorizationHeader
	}

	token, err := jwt.ParseWithClaims(headerParts[1], &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrTokenUnverifiable
		}
		return []byte(a.config.SecretJWTKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || !token.Valid {
		return 0, nil
	}

	uid, err := strconv.Atoi(claims.ID)
	if err != nil {
		return 0, CannotParseUID
	}
	return uid, nil
}

func (a *authVerificator) AddUIDToHeader(header http.Header, user models.User) error {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ID:        strconv.Itoa(user.UID),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(jwtExpire)),
	})

	jwtSignedValue, err := token.SignedString([]byte(a.config.SecretJWTKey))
	if err != nil {
		return err
	}

	header.Set(authorizationHeader, authorizationJWTPrefix+" "+jwtSignedValue)
	return nil
}

func (a *authVerificator) DeleteAuth(header http.Header) {
	header.Set(authorizationHeader, "")
}

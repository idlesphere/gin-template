package auth

import (
	env "gin-template/config"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type JWT struct {
	SigningKey []byte
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Init JWT
// Get the SALT from env
func New() *JWT {
	return &JWT{SigningKey: []byte(env.Get("SALT"))}
}

// Token Create & Refresh
func (j *JWT) Generate(username string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(8 * time.Hour)

	c := Claims{
		username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "GenchSusu",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(j.SigningKey)
}

// Token Parse
func (j *JWT) Parse(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	if token != nil {
		if c, ok := token.Claims.(*Claims); ok && token.Valid {
			return c, nil
		}
	}

	return nil, err
}

package helpers

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWTResponse model
type JWTResponse struct {
	Token     string    `json:"token"`
	ExpiresIn time.Time `json:"expires_in"`
}

// CustomClaimUser model
type CustomClaimUser struct {
	ID   uint   `json:"id"`
	jwt.StandardClaims
}

// GenerateUserToken generate token for user
func GenerateUserToken(id uint) (*JWTResponse, error) {
	var mySigningKey = []byte("sangatrahasia")
	var expiresIn = 68400
	timein := time.Now().Local().Add(time.Second * time.Duration(expiresIn))

	claim := CustomClaimUser{
		id,
		jwt.StandardClaims{
			ExpiresAt: timein.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return nil, err
	}

	// return token
	Result := JWTResponse{
		Token:     tokenString,
		ExpiresIn: timein,
	}

	return &Result, nil
}

package managers

import (
	"errors"
	"strings"
	"time"

	"handler/function/types"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("my_secret_key")

func SignAccessToken(TokenData types.TokenData) (types.TokenDTO, error) {
	return signToken(TokenData, 30, "accessToken")
}

func SignRefreshToken(TokenData types.TokenData) (types.TokenDTO, error) {
	return signToken(TokenData, 120, "refreshToken")
}

// sign jwt token
func signToken(TokenData types.TokenData, expiration time.Duration, tokenType string) (types.TokenDTO, error) {
	exp := time.Now().Add(expiration * time.Minute)
	// set token data
	claims := &jwt.MapClaims{
		"IssuedAt":  time.Now().Unix(),
		"ExpiresAt": exp.Unix(),
		"data": map[string]string{
			"sessionId": TokenData.SessionId,
			"username":  TokenData.Username,
			"type":      tokenType,
		},
	}
	// sign token
	tokenDTO := types.TokenDTO{}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtKey)
	if err != nil {
		return types.TokenDTO{}, err
	}

	tokenDTO.Value = token
	tokenDTO.Expiration = expiration.Milliseconds()
	return tokenDTO, nil
}

func VerifyAccessToken(tokenString string) (types.TokenData, error) {
	return verifyToken(tokenString, "accessToken")
}

func VerifyRefreshToken(tokenString string) (types.TokenData, error) {
	return verifyToken(tokenString, "refreshToken")
}

// verify token and return decoded TokenData
func verifyToken(tokenString string, tokenType string) (types.TokenData, error) {
	TokenData := types.TokenData{}
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		return TokenData, err
	}
	data := claims["data"].(map[string]interface{})
	TokenData.SessionId = data["sessionId"].(string)
	TokenData.Username = data["username"].(string)
	tType := data["type"].(string)
	if strings.Compare(tType, tokenType) == 0 {
		return TokenData, nil
	} else {
		return TokenData, errors.New("token verification error")
	}
}

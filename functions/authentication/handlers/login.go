package handlers

import (
	"encoding/json"
	"handler/function/managers"
	"handler/function/types"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

// handle /login request
func LoginHanlder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		LoginRequestDTO := types.LoginRequestDTO{}
		LoginResponseDTO := types.LoginResponseDTO{}
		// get payload
		body, err := io.ReadAll(r.Body)
		// check body
		if err != nil {
			managers.ThrowInternalServerError(w, err)
			return
		}
		// decode body
		err = json.Unmarshal([]byte(body), &LoginRequestDTO)
		if err != nil {
			managers.ThrowBadRequest(w, err)
			return
		}
		//TODO: do login
		// create token
		TokenData := types.TokenData{}
		TokenData.SessionId = uuid.New().String()
		TokenData.Username = LoginRequestDTO.Username
		// sign access token
		accessToken, err := managers.SignAccessToken(TokenData)
		if err != nil {
			managers.ThrowUnauthorize(w, err)
			return
		}
		// sign refresh token
		refreshToken, err := managers.SignRefreshToken(TokenData)
		if err != nil {
			managers.ThrowUnauthorize(w, err)
			return
		}
		// split tokens
		accessTokenSplitted := strings.Split(accessToken.Value, ".")
		refreshTokenSplitted := strings.Split(refreshToken.Value, ".")

		// create response
		LoginResponseDTO.AccessToken = types.TokenDTO{
			Value:      accessTokenSplitted[0] + "." + accessTokenSplitted[1],
			Expiration: accessToken.Expiration,
		}
		LoginResponseDTO.RefreshToken = types.TokenDTO{
			Value:      refreshTokenSplitted[0] + "." + refreshTokenSplitted[1],
			Expiration: refreshToken.Expiration,
		}
		LoginResponseDTO.Message = "Login successfull"

		// set token signature as cookie http only
		w.Header().Add("Set-Cookie", "accessTokenSignature="+accessTokenSplitted[2]+";Expire="+time.Unix(accessToken.Expiration, 0).String()+";HttpOnly")
		w.Header().Add("Set-Cookie", "refreshTokenSignature="+refreshTokenSplitted[2]+";Expire="+time.Unix(refreshToken.Expiration, 0).String()+";HttpOnly")

		res, err := json.Marshal(LoginResponseDTO)
		if err != nil {
			managers.ThrowInternalServerError(w, err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	}
}

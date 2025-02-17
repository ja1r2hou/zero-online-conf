package util

import "github.com/golang-jwt/jwt/v4"

func AddToken(secretKey string, iat, seconds, uid int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["uid"] = uid
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
func GetToken(secretKey string, iat, seconds, uid int64, accessToken string) (bool, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["uid"] = uid
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	tk, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return false, err
	}
	if accessToken == tk {
		return true, nil
	}
	return false, nil
}

package util

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJwtToken(studentId string, ioa, exp int64, alreadyTotp bool) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"student_id":   studentId,
		"ioa":          ioa,
		"exp":          exp,
		"already_totp": alreadyTotp,
	})
	jwtToken, err := token.SignedString([]byte(JWT_SECRET_KEY))
	if err != nil {
		return "", err
	}

	sha256Token := Sha256(jwtToken)

	return jwtToken + sha256Token, nil
}

func VerifyJwtToken(jwtToken string) (bool, string, error) {
	if len(jwtToken) < 64 {
		return false, "", errors.New("invalid token format")
	}

	token := jwtToken[:len(jwtToken)-64]
	shaSum := jwtToken[len(jwtToken)-64:]

	calculatedSha := Sha256(token)
	if calculatedSha != shaSum {
		return false, "", errors.New("token integrity check failed")
	}

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWT_SECRET_KEY), nil
	})
	if err != nil {
		return false, "", err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return false, "", errors.New("invalid token")
	}

	exp, ok := claims["exp"].(float64)
	if !ok {
		return false, "", errors.New("missing expiration time")
	}

	if int64(exp) < time.Now().Unix() {
		return false, "", errors.New("token expired")
	}

	return true, claims["student_id"].(string), nil
}

func Sha256(data string) string {
	hash := sha256.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}

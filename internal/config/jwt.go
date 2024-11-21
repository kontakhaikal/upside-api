package config

import (
	"github.com/fkrhykal/upside-api/internal/dto"
	"github.com/fkrhykal/upside-api/internal/util"
	"github.com/golang-jwt/jwt/v5"
)


type JwtCredentialUtil struct {
	key []byte
}

func (j *JwtCredentialUtil) GenerateToken(credential *dto.UserCredential) (dto.CredentialToken, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": credential.ID,
	})

	token, err := jwtToken.SignedString(j.key)

	if err != nil {
		return "", err
	}

	return dto.CredentialToken(token), nil
}

func (j *JwtCredentialUtil) VerifyToken(token dto.CredentialToken) error {
	return nil
}

func NewJwtCredentialUtil(key []byte) util.CredentialUtil {
	return &JwtCredentialUtil{
		key,
	}
}
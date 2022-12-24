package services

import (
	"crypto/rsa"
	"log"
	"macus/pkg/common"
	"macus/pkg/models"
	"macus/pkg/models/apperrors"
)

type TSConfig struct {
	PrivKey          *rsa.PrivateKey
	PubKey           *rsa.PublicKey
	IDExpirationSecs int64
}

type TokenService interface {
	NewToken(u *models.LogIn) (*models.TokenPair, error)
}

type tokenService struct {
	PrivKey          *rsa.PrivateKey
	PubKey           *rsa.PublicKey
	IDExpirationSecs int64
}

func NewTokenService(c *TSConfig) *tokenService {
	return &tokenService{
		PrivKey:          c.PrivKey,
		PubKey:           c.PubKey,
		IDExpirationSecs: c.IDExpirationSecs,
	}
}

func (s *tokenService) NewToken(u *models.LogIn) (*models.TokenPair, error) {
	idToken, err := common.GenerateIDToken(&u.UserName, s.PrivKey, s.IDExpirationSecs)
	if err != nil {
		log.Printf("Error generating idToken for uname: %v. Error: %v\n", u.UserName, err.Error())
		return nil, apperrors.NewInternal()
	}

	return &models.TokenPair{
		IDToken: models.IDToken{SS: idToken},
	}, nil
}

// https://kentakodashima.medium.com/generate-pem-keys-with-openssl-on-macos-ecac55791373
